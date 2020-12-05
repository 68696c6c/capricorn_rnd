package cmd

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/http"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func buildServer(pkg golang.IPackage, o config.CmdOptions, projectName string, a *app.App, h *http.Http) *golang.Var {
	file := pkg.AddGoFile(o.ServerFileName)
	result := makeServerVar(o, projectName, a, h)
	file.AddVar(result)
	return result
}

func makeServerVar(o config.CmdOptions, projectName string, a *app.App, h *http.Http) *golang.Var {
	return makeCommandVar(commandFuncMeta{
		name:    utils.Pascal(o.ServerFileName),
		use:     utils.Kebob(o.ServerFileName),
		short:   fmt.Sprintf("Runs the %s web server", projectName),
		long:    "",
		example: "",
		runFunc: makeServerRunFunc(o, a, h),
	})
}

func makeServerRunFunc(o config.CmdOptions, a *app.App, h *http.Http) *golang.Function {
	result := golang.NewFunction("")
	t := `
		goat.Init()

		logger := goat.GetLogger()

		db, err := goat.GetMainDB()
		if err != nil {
			goat.ExitError(errors.Wrap(err, "failed to initialize database connection"))
		}

		services, err := {{ .AppInitFuncRef }}(db, logger)
		if err != nil {
			goat.ExitError(errors.Wrap(err, "failed to initialize application services"))
		}

		r := {{ .RouterInitFuncRef }}(services)
		r.Run()
	`
	result.AddArg(o.CmdArgName, goat.MakeTypeCobraCommand())
	result.AddArg(o.ArgsArgName, golang.MakeTypeStringSlice(false))

	result.SetBodyTemplate(t, struct {
		AppInitFuncRef    string
		RouterInitFuncRef string
	}{
		AppInitFuncRef:    a.GetContainerConstructor().GetReference(),
		RouterInitFuncRef: h.GetInitRouter().GetReference(),
	})

	result.AddImportsApp(a.GetImport(), h.GetImport())
	result.AddImportsVendor(goat.ImportGoat, goat.ImportErrors, goat.ImportCobra)

	return result
}
