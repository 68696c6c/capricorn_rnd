package db

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/db/migrations"
)

func Build(pkg *golang.Package, a *app.App) {
	pkgDb := pkg.AddPackage("db")
	migrations.Build(pkgDb, a)
}
