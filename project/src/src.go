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

type Src struct {
	app  *app.App
	cmd  *cmd.Commands
	db   *db.Db
	http *http.Http
}

func Build(root utils.Directory, p *config.Project, o config.SrcOptions) *Src {
	dirSrc := utils.NewFolder(root.GetFullPath(), o.PkgName)
	srcPath := dirSrc.GetFullPath()

	pkgApp := app.NewApp(srcPath, p, o.App)
	dirSrc.AddDirectory(pkgApp)
	domainMap := pkgApp.GetDomains().GetMap()

	pkgDb := db.Build(srcPath, p.Module, o.Db, domainMap)
	dirSrc.AddDirectory(pkgDb)

	pkgHttp := http.Build(srcPath, p.Module, o.Http, pkgApp) // @TODO: of pkgApp, only using domains, error handler; seems to be mostly container stuff and domains
	dirSrc.AddDirectory(pkgHttp)

	pkgCmd := cmd.Build(srcPath, p, o.Cmd, pkgApp, pkgHttp, pkgDb.GetImportMigrations(true)) // @TODO: of pkgApp, only using container, router; func names and imports
	dirSrc.AddDirectory(pkgCmd)

	mainFile := buildMainGo(srcPath, p, o.Cmd, pkgCmd) // @TODO: of pkgApp, only using cmd.Commands.Render
	dirSrc.AddRenderableFile(mainFile)

	root.AddDirectory(dirSrc)

	return &Src{
		app:  pkgApp,
		cmd:  pkgCmd,
		db:   pkgDb,
		http: pkgHttp,
	}
}

func (s *Src) GetApp() *app.App {
	return s.app
}

func (s *Src) GetDb() *db.Db {
	return s.db
}

func buildMainGo(rootPath string, p *config.Project, o config.CmdOptions, commands *cmd.Commands) *golang.File {
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
	mainFunc.AddImportsVendor(goat.ImportSqlDriver, goat.ImportCobra, goat.ImportViper)

	pkgMain := golang.NewPackage("main", rootPath, p.Module)
	mainFile := golang.NewFile(rootPath, "main")
	mainFile.PKG = pkgMain
	mainFile.AddFunction(mainFunc)

	return mainFile
}
