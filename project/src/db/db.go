package db

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/db/migrations"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Db struct {
	*golang.Package
	pkgMigrations golang.IPackage
}

func Build(root utils.Directory, module string, o config.DbOptions, a *app.App) *Db {
	pkgDb := golang.NewPackage(o.PkgName, root.GetFullPath(), module)
	return &Db{
		Package:       pkgDb,
		pkgMigrations: migrations.Build(pkgDb, o.Migrations, a),
	}
}

func (d *Db) GetImportMigrations(underscore bool) string {
	imp := d.pkgMigrations.GetImport()
	if underscore {
		t := `// imported for auto migration registration
	_ "%s"`
		return fmt.Sprintf(t, imp)
	}
	return imp
}
