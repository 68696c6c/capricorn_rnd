package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

func makeGetFilteredQuery(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.FilterQueryFuncName)
	t := `
	result, err := {{ .QueryArgName }}.ApplyToGorm({{ .BaseQueryFuncCall }})
	if err != nil {
		return result, err
	}
	return result, nil
`

	method.AddArg(o.QueryArgName, meta.queryType)

	method.AddReturn("", meta.dbType)
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		BaseQueryFuncCall string
		QueryArgName      string
	}{
		BaseQueryFuncCall: fmt.Sprintf("%s.%s()", meta.receiverName, o.BaseQueryFuncName),
		QueryArgName:      o.QueryArgName,
	})

	method.AddImportsVendor(meta.dbType.Import)

	return method
}
