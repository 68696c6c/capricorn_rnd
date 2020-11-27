package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeConstructor(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction("New" + meta.repoInterfaceType.GetName())
	t := `
	return {{ .StructName }}{
		{{ .DbFieldName }}: {{ .DbArgName }},
	}
`
	method.AddArg(o.DbArgName, meta.dbType)

	method.AddReturn("", meta.repoInterfaceType)

	method.SetBodyTemplate(t, struct {
		StructName  string
		DbFieldName string
		DbArgName   string
	}{
		StructName:  meta.repoStructType.GetName(),
		DbFieldName: o.DbFieldName,
		DbArgName:   o.DbArgName,
	})

	method.AddImportsVendor(goat.ImportGorm)

	return method
}
