package migrations

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
)

func Build(pkg *golang.Package, o config.MigrationsOptions, domainMap domain.Map) golang.IPackage {
	pkgMigrations := pkg.AddPackage(o.PkgName)
	fileName := fmt.Sprintf("%s_%s", o.InitialMigrationTimestamp, o.InitialMigrationName)
	file := pkgMigrations.AddGoFile(fileName)

	upFuncName := fmt.Sprintf("Up%s", o.InitialMigrationTimestamp)
	downFuncName := fmt.Sprintf("Down%s", o.InitialMigrationTimestamp)

	file.AddFunction(buildInitFunc(upFuncName, downFuncName))

	file.AddFunction(buildFunc(upFuncName, domainMap))

	file.AddFunction(buildFunc(downFuncName, domainMap))

	return pkgMigrations
}

func buildInitFunc(upFuncName, downFuncName string) *golang.Function {
	result := golang.NewFunction("init")
	t := `
	goose.AddMigration({{ .UpFuncName }}, {{ .DownFuncName }})
`
	result.SetBodyTemplate(t, struct {
		UpFuncName   string
		DownFuncName string
	}{
		UpFuncName:   upFuncName,
		DownFuncName: downFuncName,
	})
	return result
}

func buildFunc(name string, domainMap domain.Map) *golang.Function {
	result := golang.NewFunction(name)
	t := `
	goat.Init()

	{{ .DbVarName }}, err := goat.GetMigrationDB()
	if err != nil {
		return errors.Wrap(err, "failed to initialize migration connection")
	}

{{ .Migrations }}

	return nil
`

	dbVarName := "db"
	var migrations []string
	for _, d := range domainMap {
		domainModel := d.GetModelType()
		migrations = append(migrations, fmt.Sprintf("	%s.AutoMigrate(&%s{})", dbVarName, domainModel.GetReference()))
		result.AddImportsApp(domainModel.GetImport())
	}

	txType := golang.MakeTypeSqlTransaction(true)
	result.AddArg("tx", txType)

	result.AddReturn("", golang.MakeTypeError())

	result.SetBodyTemplate(t, struct {
		DbVarName  string
		Migrations string
	}{
		DbVarName:  dbVarName,
		Migrations: strings.Join(migrations, "\n"),
	})

	result.AddImportsStandard(txType.GetImport())

	result.AddImportsVendor(goat.ImportGoat, goat.ImportGoose, goat.ImportErrors)

	return result
}
