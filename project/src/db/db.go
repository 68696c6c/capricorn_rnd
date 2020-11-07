package db

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/db/migrations"
)

type DB struct {
	pkg        *golang.Package
	migrations migrations.Migrations
}

func NewDB(pkg *golang.Package, a app.App) DB {
	pkgDb := pkg.AddPackage("db")
	return DB{
		pkg:        pkgDb,
		migrations: migrations.NewMigrations(pkgDb, a),
	}
}
