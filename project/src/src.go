package src

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/cmd"
	"github.com/68696c6c/capricorn_rnd/project/src/db"
	"github.com/68696c6c/capricorn_rnd/project/src/http"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func Build(root utils.Directory, p *config.Project, o config.SrcOptions) {
	pkgSrc := golang.NewPackage(o.PkgName, root.GetFullPath(), p.Module)
	srcApp := app.NewApp(pkgSrc, p, o.App)

	db.Build(pkgSrc, o.Db, srcApp)
	srcHttp := http.Build(pkgSrc, o.Http, srcApp)
	srcCommands := cmd.Build(pkgSrc, p, o.Cmd, srcApp, srcHttp)
	buildMainGo(pkgSrc, p, o.Cmd, srcCommands)

	root.AddDirectory(pkgSrc)
}

func buildMainGo(pkg golang.IPackage, p *config.Project, o config.CmdOptions, commands *cmd.Commands) {
	file := pkg.AddGoFile("main")
	mainFunc := golang.NewFunction("main")
	t := `
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("author", "{{ .AuthorName }} <{{ .AuthorEmail }}>")
	viper.SetDefault("license", "{{ .License }}")

	rootCmd := &cobra.Command{
		Use:   "{{ .RootCommandUsage }}",
		Short: "Root command for {{ .ProjectName }}",
	}

	rootCmd.SetOut(os.Stdout)
	rootCmd.AddCommand({{ .Commands.Render }}
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
`
	mainFunc.SetBodyTemplate(t, struct {
		AuthorName       string
		AuthorEmail      string
		License          string
		RootCommandUsage string
		ProjectName      string
		Commands         utils.Renderable
	}{
		AuthorName:       p.Author.Name,
		AuthorEmail:      p.Author.Email,
		License:          p.License,
		RootCommandUsage: o.RootCommandUse,
		ProjectName:      p.Name,
		Commands:         commands,
	})
	mainFunc.AddImportsStandard("os", "strings")
	mainFunc.AddImportsApp(commands.GetImport())
	mainFunc.AddImportsVendor(goat.ImportCobra, goat.ImportViper)
	file.AddFunction(mainFunc)
}
