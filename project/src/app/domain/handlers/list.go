package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeList(meta handlerMeta) *Handler {
	name := fmt.Sprintf("List%s", utils.Pascal(meta.PluralName))
	body := `
		q := query.NewQueryBuilder({{ .ContextArgName }})

		result, err := {{ .RepoRef }}.{{ .FilterFuncName }}(q)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to get {{ .PluralName }}", goat.RespondServerError)
			return
		}

		err = {{ .RepoRef }}.{{ .PageFuncName }}(q)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to count {{ .PluralName }}", goat.RespondServerError)
			return
		}

		goat.RespondData({{ .ContextArgName }}, {{ .ListResponseName }}{result, q.Pagination})
	`
	data := struct {
		ContextArgName   string
		ErrorsRef        string
		RepoRef          string
		FilterFuncName   string
		PageFuncName     string
		PluralName       string
		ListResponseName string
	}{
		ContextArgName:   meta.ContextArg.Name,
		ErrorsRef:        meta.ErrorsArg.Name,
		RepoRef:          meta.RepoArg.Name,
		FilterFuncName:   meta.RepoFilterFuncName,
		PageFuncName:     meta.RepoPageFuncName,
		PluralName:       meta.PluralName,
		ListResponseName: meta.ListResponseType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.ContextArg)

	handler.AddArgV(meta.ErrorsArg)
	handler.AddArgV(meta.RepoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbGet,
		uri:           `""`,
		requestStruct: nil,
	}
}
