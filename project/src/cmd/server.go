package cmd

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func buildServer(pkg golang.IPackage, meta *config.CmdMeta) {
	file := pkg.AddGoFile(meta.ServerFileName)
	file.AddFunction(makeServerFunc(meta))
}

func makeServerFunc(meta *config.CmdMeta) *golang.Function {
	return makeCommandFunc(commandFuncMeta{
		rootVarName: meta.RootVarName,
		use:         utils.Kebob(meta.ServerFileName),
		short:       fmt.Sprintf("Runs the %s web server", meta.ProjectName),
		long:        "",
		example:     "",
		runFunc:     makeServerRunFunc(config.AppInitFuncName, config.RouterInitFuncName),
	})
}

func makeServerRunFunc(appInitFuncName, routerInitFuncName string) *golang.Function {
	result := golang.NewFunction("")
	t := `
			goat.Init()

			logger := goat.GetLogger()

			db, err := goat.GetMainDB()
			if err != nil {
				goat.ExitError(errors.Wrap(err, "failed to initialize database connection"))
			}

			services, err := {{ .AppInitFuncName }}(db, logger)
			if err != nil {
				goat.ExitError(errors.Wrap(err, "failed to initialize application services"))
			}

			{{ .RouterInitFuncName }}(services)
		`
	result.AddArg("cmd", goat.MakeTypeCobraCommand())
	result.AddArg("args", golang.MakeTypeStringSlice(false))
	result.SetBodyTemplate(t, struct {
		AppInitFuncName    string
		RouterInitFuncName string
	}{
		AppInitFuncName:    appInitFuncName,
		RouterInitFuncName: routerInitFuncName,
	})
	result.AddImportsVendor(goat.ImportGoat, goat.ImportErrors)
	return result
}
