package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeDelete(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.DeleteFuncName)
	t := `
	errs :=  {{ .DbRef }}.Delete({{ .ModelArgName }}).GetErrors()
	if len(errs) > 0 {
		return goat.ErrorsToError(errs)
	}
	return nil
`

	argType := meta.modelType.CopyType()
	argType.IsPointer = true
	method.AddArg(o.ModelArgName, argType)

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		DbRef        string
		ModelArgName string
	}{
		DbRef:        meta.dbFieldRef,
		ModelArgName: o.ModelArgName,
	})

	method.AddImportsVendor(goat.ImportGoat, goat.ImportErrors)

	return method
}
