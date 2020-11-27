package container

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeConstructor(o config.ServiceContainerOptions, containerType *golang.Struct, domains domain.Map) *golang.Function {
	method := golang.NewFunction(o.AppInitFuncName)
	t := `
	if {{ .SingletonName }} != ({{ .TypeName }}{}) {
		return {{ .SingletonName }}, nil
	}
{{ .Declarations }}

	container = {{ .TypeName }}{
		{{ .DbFieldName }}:     {{ .DbArgName }},
		{{ .LoggerFieldName }}: {{ .LoggerArgName }},
		{{ .ErrorsFieldName }}: goat.NewErrorHandler({{ .LoggerArgName }}),
{{ .Fields }}
	}

	return {{ .SingletonName }}, nil
`
	method.AddArg(o.DbArgName, goat.MakeTypeDbConnection())

	method.AddArg(o.LoggerArgName, goat.MakeTypeLogger())

	method.AddReturn("", containerType)

	method.AddReturn("", golang.MakeTypeError())

	var declarations []string
	var fields []string
	for resourceName, d := range domains {
		if !d.HasRepo() {
			continue
		}

		method.AddImportsApp(d.GetImport())

		repoConstructor := d.GetRepoConstructor()
		repoFieldName := utils.Pascal(d.GetExternalRepoName())
		if d.HasService() {
			varName := utils.Camel(o.RepoVarNameTemplate.Parse(resourceName))

			repoDec := fmt.Sprintf("%s := %s(%s)", varName, repoConstructor.GetReference(), o.DbArgName)
			declarations = append(declarations, "\t"+repoDec)

			fields = append(fields, fmt.Sprintf("\t\t%s: %s,", repoFieldName, varName))

			serviceConstructor := d.GetServiceConstructor()
			fields = append(fields, fmt.Sprintf("\t\t%s: %s(%s),", utils.Pascal(d.GetExternalServiceName()), serviceConstructor.GetReference(), varName))
		} else {
			fields = append(fields, fmt.Sprintf("\t\t%s: %s(%s),", repoFieldName, repoConstructor.GetReference(), o.DbArgName))
		}
	}

	method.SetBodyTemplate(t, struct {
		SingletonName   string
		TypeName        string
		DbFieldName     string
		DbArgName       string
		LoggerFieldName string
		LoggerArgName   string
		ErrorsFieldName string
		Declarations    string
		Fields          string
	}{
		SingletonName:   o.SingletonName,
		TypeName:        containerType.GetName(),
		DbFieldName:     o.DbFieldName,
		DbArgName:       o.DbArgName,
		LoggerFieldName: o.LoggerFieldName,
		LoggerArgName:   o.LoggerArgName,
		ErrorsFieldName: o.ErrorsFieldName,
		Declarations:    strings.Join(declarations, "\n"),
		Fields:          strings.Join(fields, "\n"),
	})

	method.AddImportsVendor(goat.ImportGoat, goat.ImportGorm, goat.ImportLogrus)

	return method
}
