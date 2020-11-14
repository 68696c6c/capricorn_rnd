package repo

import "github.com/68696c6c/capricorn_rnd/golang"

func makeSave(meta methodMeta) *golang.Function {
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

	method.AddArg(meta.modelArgName, meta.modelType)

	method.AddReturn("", golang.MakeErrorType())

	method.SetBodyTemplate(t, struct {
		DbRef        string
		ModelArgName string
	}{
		DbRef:        meta.dbFieldRef,
		ModelArgName: meta.modelArgName,
	})

	method.AddImportsVendor(golang.ImportGoat)

	return method
}
