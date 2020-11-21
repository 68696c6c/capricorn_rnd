package repo

import (
	"fmt"
	"github.com/68696c6c/capricorn_rnd/project/goat"

	"github.com/68696c6c/capricorn_rnd/golang"
)

func makeFilter(meta methodMeta) *golang.Function {
	method := golang.NewFunction(meta.filterFuncName)
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
		FilterQueryFuncCall: fmt.Sprintf("%s.%s(%s)", meta.receiverName, meta.filterQueryFuncName, meta.queryArgName),
		PageQueryFuncCall:   fmt.Sprintf("%s.%s(%s)", meta.receiverName, meta.pageQueryFuncName, meta.queryArgName),
	})

	method.AddImportsVendor(goat.ImportGoat, meta.queryType.Import, goat.ImportErrors)

	return method
}
