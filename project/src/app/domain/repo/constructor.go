package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeConstructor(meta methodMeta) *golang.Function {
	repoName := meta.repoStructType.Name
	method := golang.NewFunction("New" + repoName)
	t := `
	return {{ .StructName }}{
		{{ .DbFieldName }}: {{ .DbArgName }},
	}
`
	dbArgName := "dbConnection"
	method.AddArg(dbArgName, meta.dbType)

	method.AddReturn("", meta.repoInterfaceType)

	method.SetBodyTemplate(t, struct {
		StructName  string
		DbFieldName string
		DbArgName   string
	}{
		StructName:  repoName,
		DbFieldName: meta.dbFieldName,
		DbArgName:   dbArgName,
	})

	method.AddImportsVendor(goat.ImportGorm)

	return method
}
