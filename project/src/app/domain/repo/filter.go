package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeFilter(o config.RepoOptions, meta *methodMeta) *golang.Function {
	method := golang.NewFunction(o.FilterFuncName)
	t := `
	dataQuery, err := {{ .FilterQueryFuncCall }}
	if err != nil {
		return result, errors.Wrap(err, "failed to build filter {{ .PluralName }} query")
	}

	errs := dataQuery.Find(&result).GetErrors()
	if len(errs) > 0 && goat.ErrorsBesidesRecordNotFound(errs) {
		err := goat.ErrorsToError(errs)
		return result, errors.Wrap(err, "failed to execute filter {{ .PluralName }} data query")
	}

	err = {{ .PageQueryFuncCall }}
	if err != nil {
		return result, err
	}

	return result, nil
`

	method.AddArg(o.QueryArgName, meta.queryType)

	returnType := meta.modelType.CopyType()
	returnType.IsPointer = true
	method.AddReturn("result", golang.MakeSliceType(false, returnType))
	method.AddReturn("err", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		PluralName          string
		FilterQueryFuncCall string
		PageQueryFuncCall   string
	}{
		PluralName:          meta.modelPlural,
		FilterQueryFuncCall: fmt.Sprintf("%s.%s(%s)", meta.receiverName, o.FilterQueryFuncName, o.QueryArgName),
		PageQueryFuncCall:   fmt.Sprintf("%s.%s(%s)", meta.receiverName, o.PaginationFuncName, o.QueryArgName),
	})

	method.AddImportsVendor(goat.ImportGoat, meta.queryType.Import, goat.ImportErrors)

	return method
}
