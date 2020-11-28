package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

func makeGetBaseQuery(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.BaseQueryFuncName)
	t := `
	return {{ .DbRef }}.Model(&{{ .ModelTypeName }}{})
`

	method.AddReturn("", meta.dbType)

	method.SetBodyTemplate(t, struct {
		DbRef         string
		ModelTypeName string
	}{
		DbRef:         meta.dbFieldRef,
		ModelTypeName: meta.modelType.CopyType().GetReference(),
	})

	method.AddImportsVendor(meta.dbType.Import)

	return method
}
