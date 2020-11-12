package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
)

func makeFilter(meta methodMeta) *golang.Function {
	method := golang.NewFunction(meta.baseImport, meta.pkgName, "Filter")
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

	method.AddArg(meta.queryArgName, meta.queryType)

	method.AddReturn("result", golang.MakeSliceType(false, meta.modelType))
	method.AddReturn("err", golang.MakeErrorType())

	method.SetBodyTemplate(t, struct {
		PluralName          string
		FilterQueryFuncCall string
		PageQueryFuncCall   string
	}{
		PluralName:          meta.modelPlural,
		FilterQueryFuncCall: fmt.Sprintf("%s.%s(%s)", meta.receiverName, meta.filterQueryFuncName, meta.queryArgName),
		PageQueryFuncCall:   fmt.Sprintf("%s.%s(%s)", meta.receiverName, meta.pageQueryFuncName, meta.queryArgName),
	})

	method.AddImportsVendor(golang.ImportGoat, meta.queryType.Import, golang.ImportErrors)

	return method
}
