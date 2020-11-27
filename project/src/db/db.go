package db

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/db/migrations"
)

func Build(pkg golang.IPackage, o config.DbOptions, a *app.App) {
	pkgDb := pkg.AddPackage(o.PkgName)
	migrations.Build(pkgDb, o.Migrations, a)
}
