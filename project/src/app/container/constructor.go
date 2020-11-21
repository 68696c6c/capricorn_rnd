package container

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeConstructor(meta containerMeta) *golang.Function {
	method := golang.NewFunction("GetApp")
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
	dbArgName := "dbConnection"
	method.AddArg(dbArgName, goat.MakeTypeDbConnection())

	loggerArgName := "logger"
	method.AddArg(loggerArgName, goat.MakeTypeLogger())

	method.AddReturn("", meta.structType)

	method.AddReturn("", golang.MakeTypeError())

	var declarations []string
	var fields []string
	for key, d := range meta.domains {
		method.AddImportsApp(d.GetImport())

		repoConstructor := d.Repo.GetConstructor()
		repoFieldName := utils.Pascal(d.GetExternalRepoName())
		if d.Service != nil {
			varName := utils.Camel(key + "_repo")

			repoDec := fmt.Sprintf("%s := %s(%s)", varName, repoConstructor.GetReference(), dbArgName)
			declarations = append(declarations, "\t"+repoDec)

			fields = append(fields, fmt.Sprintf("\t\t%s: %s,", repoFieldName, varName))

			serviceConstructor := d.Service.GetConstructor()
			fields = append(fields, fmt.Sprintf("\t\t%s: %s(%s),", utils.Pascal(d.GetExternalServiceName()), serviceConstructor.GetReference(), varName))
		} else {
			fields = append(fields, fmt.Sprintf("\t\t%s: %s(%s),", repoFieldName, repoConstructor.GetReference(), dbArgName))
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
		SingletonName:   meta.singletonName,
		TypeName:        meta.structType.GetName(),
		DbFieldName:     meta.dbFieldName,
		DbArgName:       dbArgName,
		LoggerFieldName: meta.loggerFieldName,
		LoggerArgName:   loggerArgName,
		ErrorsFieldName: meta.errorsFieldName,
		Declarations:    strings.Join(declarations, "\n"),
		Fields:          strings.Join(fields, "\n"),
	})

	method.AddImportsVendor(goat.ImportGoat, goat.ImportGorm, goat.ImportLogrus)

	return method
}
