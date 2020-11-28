package cmd

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func buildMigrate(pkg golang.IPackage, o config.CmdOptions, migrationsImport string) *golang.Var {
	file := pkg.AddGoFile(o.MigrateFileName)
	file.AddImportsApp(migrationsImport)

	result := makeMigrateVar(o)
	file.AddVar(result)

	return result
}

func makeMigrateVar(o config.CmdOptions) *golang.Var {
	t := `{{ .RootUse }} {{ .MigrateUse }} status
{{ .RootUse }} {{ .MigrateUse }} create init sql
{{ .RootUse }} {{ .MigrateUse }} create add_some_column sql
{{ .RootUse }} {{ .MigrateUse }} create fetch_user_data go
{{ .RootUse }} {{ .MigrateUse }} up`
	example, err := utils.ParseTemplate("template_migrate_example", t, struct {
		RootUse    string
		MigrateUse string
	}{
		RootUse:    o.RootCommandUse,
		MigrateUse: o.MigrateCommandUse,
	})
	if err != nil {
		panic(err)
	}
	return makeCommandVar(commandFuncMeta{
		name:    utils.Pascal(o.MigrateFileName),
		use:     fmt.Sprintf("%s %s [OPTIONS] COMMAND", o.RootCommandUse, o.MigrateCommandUse),
		short:   "Root migration command.",
		long:    "",
		example: example,
		runFunc: makeMigrateRunFunc(o),
	})
}

func makeMigrateRunFunc(o config.CmdOptions) *golang.Function {
	result := golang.NewFunction("")
	t := `
		goat.Init()

		db, err := goat.GetMainDB()
		if err != nil {
			goat.ExitError(errors.Wrap(err, "error initializing migration connection"))
		}

		err = goose.SetDialect("mysql")
		if err != nil {
			goat.ExitError(errors.Wrap(err, "error initializing goose"))
		}

		var arguments []string
		if len(args) > 1 {
			arguments = args[1:]
		}

		err = goose.Run(args[0], db.DB(), ".", arguments...)
		if err != nil {
			goat.ExitError(err)
		}

		goat.ExitSuccess()
	`
	result.AddArg(o.CmdArgName, goat.MakeTypeCobraCommand())
	result.AddArg(o.ArgsArgName, golang.MakeTypeStringSlice(false))
	result.SetBodyTemplate(t, nil)
	result.AddImportsVendor(goat.ImportGoat, goat.ImportErrors, goat.ImportGoose, goat.ImportCobra)
	return result
}
