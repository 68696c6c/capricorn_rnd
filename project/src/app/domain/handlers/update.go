package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeUpdate(o config.HandlersOptions, meta handlerMeta) *Handler {
	name := utils.Pascal(o.UpdateNameTemplate.Parse(meta.resourceName))
	body := `
		i := c.Param("{{ .IdParamName }}")
		id, err := goat.ParseID(i)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to parse id: "+i, goat.RespondBadRequestError)
			return
		}

		_, err = {{ .RepoRef }}.{{ .RepoGetByIdFuncName }}(id)
		if err != nil {
			if goat.IsNotFoundError(err) {
				{{ .ErrorsRef }}.HandleMessage({{ .ContextArgName }}, "{{ .SingleName }} not found", goat.RespondNotFoundError)
				return
			} else {
				{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to get {{ .SingleName }}", goat.RespondServerError)
				return
			}
		}

		req, ok := goat.GetRequest(c).(*{{ .UpdateRequestTypeName }})
		if !ok {
			errorHandler.HandleMessage(c, "failed to get request", goat.RespondBadRequestError)
			return
		}

		err = {{ .RepoRef }}.{{ .RepoSaveFuncName }}(&req.{{ .ModelTypeName }})
		if err != nil {
			errorHandler.HandleErrorM(c, err, "failed to save {{ .SingleName }}", goat.RespondServerError)
			return
		}

		goat.RespondData({{ .ContextArgName }}, {{ .ResourceResponseTypeName }}{req.{{ .ModelTypeName }}})
	`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		RepoGetByIdFuncName      string
		RepoSaveFuncName         string
		IdParamName              string
		SingleName               string
		ResourceResponseTypeName string
		UpdateRequestTypeName    string
		ModelTypeName            string
	}{
		ContextArgName:           meta.contextArg.Name,
		ErrorsRef:                meta.errorsArg.Name,
		RepoRef:                  meta.repoArg.Name,
		RepoGetByIdFuncName:      o.RepoGetByIdFuncName,
		RepoSaveFuncName:         o.RepoSaveFuncName,
		IdParamName:              o.ParamNameId,
		SingleName:               meta.nameSingular,
		ResourceResponseTypeName: meta.resourceResponseType.Name,
		UpdateRequestTypeName:    meta.requestUpdateType.Name,
		ModelTypeName:            meta.modelTypeName,
	}

	handler := makeHandlerFunc(name, body, data, meta.contextArg)

	handler.AddArgV(meta.errorsArg)
	handler.AddArgV(meta.repoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbPut,
		uri:           fmt.Sprintf(`"/:%s"`, o.ParamNameId),
		requestStruct: meta.requestUpdateType,
	}
}
