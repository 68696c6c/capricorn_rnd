package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeGetById(meta methodMeta) *golang.Function {
	method := golang.NewFunction("GetById")
	t := `
	{{ .ModelVarName }} := &{{ .ModelTypeName }}{
		Model: goat.Model{
			ID: id,
		},
	}
	errs := {{ .DbRef }}.First({{ .ModelVarName }}).GetErrors()
	if len(errs) > 0 {
		return {{ .ModelVarName }}, goat.ErrorsToError(errs)
	}
	return {{ .ModelVarName }}, nil
`

	idArgName := "id"
	method.AddArg(idArgName, goat.MakeTypeId())

	returnType := meta.modelType.CopyType()
	returnType.IsPointer = true
	method.AddReturn("", returnType)
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		DbRef         string
		IdArgName     string
		ModelVarName  string
		ModelTypeName string
	}{
		DbRef:         meta.dbFieldRef,
		IdArgName:     idArgName,
		ModelVarName:  meta.modelArgName,
		ModelTypeName: meta.modelType.Name,
	})

	method.AddImportsVendor(goat.ImportGoat)

	return method
}
