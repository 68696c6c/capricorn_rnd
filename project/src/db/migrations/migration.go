package migrations

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type Map map[string]Migration

type Migrations struct {
	pkg        *golang.Package
	migrations Map
}

type Migration struct {
	file *golang.File
}

func NewMigrations(pkg *golang.Package, a app.App) Migrations {
	result := make(Map)
	pkgMigrations := pkg.AddPackage("migrations")
	result["initial"] = newMigration(pkgMigrations, "initial")
	return Migrations{
		pkg:        pkgMigrations,
		migrations: result,
	}
}

func newMigration(pkg *golang.Package, name string) Migration {
	return Migration{
		file: pkg.AddGoFile(name),
	}
}
