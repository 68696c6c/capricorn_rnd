package src

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
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
	cmd.Build(pkgSrc, p, o.Cmd, srcApp, srcHttp)
	buildMainGo(pkgSrc)

	root.AddDirectory(pkgSrc)
}

type MainGo struct {
	*golang.File
}

func buildMainGo(pkg golang.IPackage) MainGo {
	return MainGo{
		File: pkg.AddGoFile("main"),
	}
}
