package service

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeConstructor(serviceStruct, serviceInterface, repoFieldType golang.IType, repoFieldName string) *golang.Function {
	method := golang.NewFunction("New" + serviceInterface.GetName())
	t := `
	return {{ .StructName }}{
		{{ .RepoFieldName }}: {{ .RepoArgName }},
	}
`
	repoArgName := "repo"
	method.AddArg(repoArgName, repoFieldType)

	method.AddReturn("", serviceInterface)

	method.SetBodyTemplate(t, struct {
		StructName    string
		RepoFieldName string
		RepoArgName   string
	}{
		StructName:    serviceStruct.GetName(),
		RepoFieldName: repoFieldName,
		RepoArgName:   repoArgName,
	})

	method.AddImportsVendor(goat.ImportGorm)

	return method
}
