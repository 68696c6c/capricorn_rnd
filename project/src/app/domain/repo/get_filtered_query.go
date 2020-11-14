package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
)

func makeGetFilteredQuery(meta methodMeta) *golang.Function {
	method := golang.NewFunction(meta.filterQueryFuncName)
	t := `
	result, err := {{ .QueryArgName }}.ApplyToGorm({{ .BaseQueryFuncCall }})
	if err != nil {
		return result, err
	}
	return result, nil
`

	method.AddArg(meta.queryArgName, meta.queryType)

	method.AddReturn("", meta.dbType)
	method.AddReturn("", golang.MakeErrorType())

	method.SetBodyTemplate(t, struct {
		BaseQueryFuncCall string
		QueryArgName      string
	}{
		BaseQueryFuncCall: fmt.Sprintf("%s.%s()", meta.receiverName, meta.baseQueryFuncName),
		QueryArgName:      meta.queryArgName,
	})

	method.AddImportsVendor(meta.dbType.Import)

	return method
}
