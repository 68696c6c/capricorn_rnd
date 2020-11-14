package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeDelete(meta methodMeta) *golang.Function {
	method := golang.NewFunction("Delete")
	t := `
	errs :=  {{ .DbRef }}.Delete({{ .ModelArgName }}).GetErrors()
	if len(errs) > 0 {
		return goat.ErrorsToError(errs)
	}
	return nil
`

	method.AddArg(meta.modelArgName, meta.modelType)

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		DbRef        string
		ModelArgName string
	}{
		DbRef:        meta.dbFieldRef,
		ModelArgName: meta.modelArgName,
	})

	method.AddImportsVendor(goat.ImportGoat, goat.ImportErrors)

	return method
}
