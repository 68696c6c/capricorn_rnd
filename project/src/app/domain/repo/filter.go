package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

func makeFilter(baseImport, pkgName, receiverName string, modelType model.Type) *golang.Function {
	filter := golang.NewFunction(baseImport, pkgName, "Filter")
	t := `
	dataQuery, err := {{ .ReceiverName }}.getFilteredQuery(q)
	if err != nil {
		return result, errors.Wrap(err, "failed to build filter sites query")
	}

	errs := dataQuery.Find(&result).GetErrors()
	if len(errs) > 0 && goat.ErrorsBesidesRecordNotFound(errs) {
		err := goat.ErrorsToError(errs)
		return result, errors.Wrap(err, "failed to execute filter sites data query")
	}

	if err := {{ .ReceiverName }}.applyPaginationToQuery(q); err != nil {
		return result, err
	}

	return result, nil
`
	mt := modelType
	mt.IsPointer = true

	queryType := golang.MakeQueryType()
	filter.AddArg("q", queryType)

	filter.AddReturn("result", golang.MakeSliceType(false, mt))
	filter.AddReturn("err", golang.MakeErrorType())

	filter.SetBodyTemplate(t, struct {
		ReceiverName string
	}{
		ReceiverName: receiverName,
	})

	filter.AddImportsVendor(golang.ImportErrors, queryType.Import)

	return filter
}
