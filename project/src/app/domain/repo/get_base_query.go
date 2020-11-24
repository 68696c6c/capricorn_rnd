package repo

import "github.com/68696c6c/capricorn_rnd/golang"

func makeGetBaseQuery(meta *methodMeta) *golang.Function {
	method := golang.NewFunction(meta.baseQueryFuncName)
	t := `
	return {{ .DbRef }}.Model(&{{ .ModelTypeName }}{})
`

	method.AddReturn("", meta.dbType)

	method.SetBodyTemplate(t, struct {
		DbRef         string
		ModelTypeName string
	}{
		DbRef:         meta.dbFieldRef,
		ModelTypeName: meta.modelType.GetName(),
	})

	method.AddImportsVendor(meta.dbType.Import)

	return method
}
