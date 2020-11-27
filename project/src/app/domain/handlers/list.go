package handlers

import (
	"github.com/68696c6c/capricorn_rnd/project/config"

	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeList(o config.HandlersOptions, meta handlerMeta) *Handler {
	name := utils.Pascal(o.ListNameTemplate.Parse(meta.resourceName))
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
		ContextArgName:   meta.contextArg.Name,
		ErrorsRef:        meta.errorsArg.Name,
		RepoRef:          meta.repoArg.Name,
		FilterFuncName:   o.RepoFilterFuncName,
		PageFuncName:     o.RepoPaginationFuncName,
		PluralName:       meta.namePlural,
		ListResponseName: meta.listResponseType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.contextArg)

	handler.AddArgV(meta.errorsArg)
	handler.AddArgV(meta.repoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbGet,
		uri:           `""`,
		requestStruct: nil,
	}
}
