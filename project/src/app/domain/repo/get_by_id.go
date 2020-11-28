package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeGetById(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.GetByIdFuncName)
	t := `
	{{ .ModelVarName }} := &{{ .ModelTypeName }}{
		Model: goat.Model{
			ID: id,
		},
	}
	errs := {{ .DbRef }}.First({{ .ModelVarName }}).GetErrors()
	if len(errs) > 0 {
		if goat.RecordNotFound(errs) {
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, goat.ErrorsToError(errs)
		}
	}
	return {{ .ModelVarName }}, nil
`

	method.AddArg(o.IdArgName, goat.MakeTypeId())

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
		IdArgName:     o.IdArgName,
		ModelVarName:  o.ModelArgName,
		ModelTypeName: meta.modelType.CopyType().GetReference(),
	})

	method.AddImportsVendor(goat.ImportGoat, goat.ImportGorm)

	return method
}
