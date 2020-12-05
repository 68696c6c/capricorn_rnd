package handlers

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeView(o config.HandlersOptions, meta handlerMeta) *Handler {
	name := utils.Pascal(o.ViewNameTemplate.Parse(meta.resourceName))
	body := `
		i := c.Param("{{ .IdParamName }}")
		id, err := goat.ParseID(i)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to parse id: "+i, goat.RespondBadRequestError)
			return
		}

		m, err := {{ .RepoRef }}.{{ .RepoGetByIdFuncName }}(id)
		if err != nil {
			if goat.IsNotFoundError(err) {
				{{ .ErrorsRef }}.HandleMessage({{ .ContextArgName }}, "{{ .SingleName }} not found", goat.RespondNotFoundError)
				return
			} else {
				{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to get {{ .SingleName }}", goat.RespondServerError)
				return
			}
		}

		goat.RespondData({{ .ContextArgName }}, {{ .ResourceResponseTypeName }}{*m})
	`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		RepoGetByIdFuncName      string
		IdParamName              string
		SingleName               string
		ResourceResponseTypeName string
	}{
		ContextArgName:           meta.contextArg.Name,
		ErrorsRef:                meta.errorsArg.Name,
		RepoRef:                  meta.repoArg.Name,
		RepoGetByIdFuncName:      o.RepoGetByIdFuncName,
		IdParamName:              o.ParamNameId,
		SingleName:               meta.nameSingular,
		ResourceResponseTypeName: meta.resourceResponseType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.contextArg)

	handler.AddArgV(meta.errorsArg)
	handler.AddArgV(meta.repoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbGet,
		uri:           `"/:%s"`,
		requestStruct: nil,
	}
}
