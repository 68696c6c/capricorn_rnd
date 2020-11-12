package repo

import "github.com/68696c6c/capricorn_rnd/golang"

func makeGetById(meta methodMeta) *golang.Function {
	method := golang.NewFunction(meta.baseImport, meta.pkgName, "GetById")
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
	method.AddArg(idArgName, golang.MakeIdType())

	method.AddReturn("", meta.modelType)
	method.AddReturn("", golang.MakeErrorType())

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

	method.AddImportsVendor(golang.ImportGoat)

	return method
}
