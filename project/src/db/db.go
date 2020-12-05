package db

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
	"github.com/68696c6c/capricorn_rnd/project/src/db/migrations"
)

type Db struct {
	*golang.Package
	pkgMigrations golang.IPackage
}

func Build(rootPath string, module string, o config.DbOptions, domainMap domain.Map) *Db {
	pkgDb := golang.NewPackage(o.PkgName, rootPath, module)
	return &Db{
		Package:       pkgDb,
		pkgMigrations: migrations.Build(pkgDb, o.Migrations, domainMap),
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
