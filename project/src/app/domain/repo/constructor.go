package repo

import "github.com/68696c6c/capricorn_rnd/golang"

func makeConstructor(meta methodMeta) *golang.Function {
	repoName := meta.repoStructType.Name
	method := golang.NewFunction(meta.baseImport, meta.pkgName, "New"+repoName)
	t := `
	return {{ .StructName }}{
		{{ .DbFieldName }}: {{ .DbArgName }},
	}
`
	dbArgName := "dbConnection"
	method.AddArg(dbArgName, meta.dbType)

	scopedInterfaceType := meta.repoInterfaceType
	scopedInterfaceType.Package = ""
	method.AddReturn("", scopedInterfaceType)

	method.SetBodyTemplate(t, struct {
		StructName  string
		DbFieldName string
		DbArgName   string
	}{
		StructName:  repoName,
		DbFieldName: meta.dbFieldName,
		DbArgName:   dbArgName,
	})

	method.AddImportsVendor(golang.ImportGorm)

	return method
}
