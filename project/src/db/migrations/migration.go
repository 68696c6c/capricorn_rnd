package migrations

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

func Build(pkg *golang.Package, o config.MigrationsOptions, a *app.App) {
	pkgMigrations := pkg.AddPackage(o.PkgName)
	pkgMigrations.AddGoFile(o.FileName)
}
