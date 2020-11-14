package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
)

func makeApplyPaginationToQuery(meta methodMeta) *golang.Function {
	method := golang.NewFunction(meta.pageQueryFuncName)
	t := `
	err := goat.ApplyPaginationToQuery({{ .QueryArgName }}, {{ .BaseQueryFuncCall }})
	if err != nil {
		return errors.Wrap(err, "failed to set {{ .PluralName }} query pagination")
	}
	return nil
`

	method.AddArg(meta.queryArgName, meta.queryType)

	method.AddReturn("", golang.MakeErrorType())

	method.SetBodyTemplate(t, struct {
		PluralName        string
		QueryArgName      string
		BaseQueryFuncCall string
	}{
		PluralName:        meta.modelPlural,
		QueryArgName:      meta.queryArgName,
		BaseQueryFuncCall: fmt.Sprintf("%s.%s()", meta.receiverName, meta.baseQueryFuncName),
	})

	method.AddImportsVendor(golang.ImportErrors)

	return method
}
