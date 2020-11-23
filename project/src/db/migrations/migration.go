package migrations

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

func Build(pkg *golang.Package, a *app.App) {
	pkgMigrations := pkg.AddPackage("migrations")
	pkgMigrations.AddGoFile("initial")
}
