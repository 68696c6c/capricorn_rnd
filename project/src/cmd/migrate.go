package cmd

import (
	"fmt"
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func buildMigrate(pkg golang.IPackage, meta *config.CmdMeta) {
	file := pkg.AddGoFile(meta.MigrateFileName)
	file.AddFunction(makeMigrateFunc(meta))
}

func makeMigrateFunc(meta *config.CmdMeta) *golang.Function {
	t := `{{ .RootUse }} {{ .MigrateUse }} status
{{ .RootUse }} {{ .MigrateUse }} create init sql
{{ .RootUse }} {{ .MigrateUse }} create add_some_column sql
{{ .RootUse }} {{ .MigrateUse }} create fetch_user_data go
{{ .RootUse }} {{ .MigrateUse }} up`
	example, err := utils.ParseTemplate(t, "template_migrate_example", struct {
		RootUse    string
		MigrateUse string
	}{
		RootUse:    meta.RootCommandUse,
		MigrateUse: meta.MigrateFileName,
	})
	if err != nil {
		panic(err)
	}
	return makeCommandFunc(commandFuncMeta{
		rootVarName: meta.RootVarName,
		use:         fmt.Sprintf("%s %s [OPTIONS] COMMAND", meta.RootCommandUse, meta.MigrateFileName),
		short:       "Root migration command.",
		long:        "",
		example:     example,
		runFunc:     makeMigrateRunFunc(config.AppInitFuncName, config.RouterInitFuncName),
	})
}

func makeMigrateRunFunc(appInitFuncName, routerInitFuncName string) *golang.Function {
	result := golang.NewFunction("")
	t := `
			goat.Init()

			db, err := goat.GetMainDB()
			if err != nil {
				goat.ExitError(errors.Wrap(err, "error initializing migration connection"))
			}

			if err := goose.SetDialect("mysql"); err != nil {
				goat.ExitError(errors.Wrap(err, "error initializing goose"))
			}

			var arguments []string
			if len(args) > 1 {
				arguments = args[1:]
			}

			if err := goose.Run(args[0], db.DB(), ".", arguments...); err != nil {
				goat.ExitError(err)
			}

			goat.ExitSuccess()
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
	result.AddImportsVendor(goat.ImportGoat, goat.ImportErrors, goat.ImportGoose, goat.ImportSqlDriver)
	return result
}
