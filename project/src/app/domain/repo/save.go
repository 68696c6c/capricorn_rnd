package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeSave(meta *methodMeta) *golang.Function {
	method := golang.NewFunction("Save")
	t := `
	var errs []error
	if {{ .ModelArgName }}.Model.ID.Valid() {
		errs = {{ .DbRef }}.Save({{ .ModelArgName }}).GetErrors()
	} else {
		errs = {{ .DbRef }}.Create({{ .ModelArgName }}).GetErrors()
	}
	if len(errs) > 0 {
		return goat.ErrorsToError(errs)
	}
	return nil
`

	modelType := meta.modelType.CopyType()
	modelType.IsPointer = true
	method.AddArg(meta.modelArgName, modelType)

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		DbRef        string
		ModelArgName string
	}{
		DbRef:        meta.dbFieldRef,
		ModelArgName: meta.modelArgName,
	})

	method.AddImportsVendor(goat.ImportGoat)

	return method
}
