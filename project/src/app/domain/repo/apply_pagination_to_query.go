package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeApplyPaginationToQuery(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.PaginationFuncName)
	t := `
	err := goat.ApplyPaginationToQuery({{ .QueryArgName }}, {{ .BaseQueryFuncCall }})
	if err != nil {
		return errors.Wrap(err, "failed to set {{ .PluralName }} query pagination")
	}
	return nil
`

	method.AddArg(o.QueryArgName, meta.queryType)

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		PluralName        string
		QueryArgName      string
		BaseQueryFuncCall string
	}{
		PluralName:        meta.modelPlural,
		QueryArgName:      o.QueryArgName,
		BaseQueryFuncCall: fmt.Sprintf("%s.%s()", meta.receiverName, o.BaseQueryFuncName),
	})

	method.AddImportsVendor(goat.ImportErrors)

	return method
}
