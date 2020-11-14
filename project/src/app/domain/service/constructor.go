package service

import "github.com/68696c6c/capricorn_rnd/golang"

func makeConstructor(serviceStruct *golang.Struct, serviceInterface *golang.Interface, repoFieldType golang.IType, repoFieldName string) *golang.Function {
	method := golang.NewFunction("New" + serviceStruct.Name)
	t := `
	return {{ .StructName }}{
		{{ .RepoFieldName }}: {{ .RepoArgName }},
	}
`
	repoArgName := "repo"
	method.AddArg(repoArgName, repoFieldType)

	scopedInterfaceType := serviceInterface
	scopedInterfaceType.Package = ""
	method.AddReturn("", scopedInterfaceType)

	method.SetBodyTemplate(t, struct {
		StructName    string
		RepoFieldName string
		RepoArgName   string
	}{
		StructName:    serviceStruct.Name,
		RepoFieldName: repoFieldName,
		RepoArgName:   repoArgName,
	})

	method.AddImportsVendor(golang.ImportGorm)

	return method
}
