package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeUpdate(meta handlerGroupMeta) *Handler {
	name := fmt.Sprintf("Update%s", utils.Pascal(meta.SingleName))
	body := `
		i := c.Param("{{ .IdParamName }}")
		id, err := goat.ParseID(i)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to parse id: "+i, goat.RespondBadRequestError)
			return
		}

		_, err = {{ .RepoRef }}.GetByID(id)
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

		err = sitesRepo.Save(&req.Site)
		if err != nil {
			errorHandler.HandleErrorM(c, err, "failed to save {{ .SingleName }}", goat.RespondServerError)
			return
		}

		goat.RespondData({{ .ContextArgName }}, {{ .ResourceResponseTypeName }}{m})
`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		IdParamName              string
		SingleName               string
		ResourceResponseTypeName string
		UpdateRequestTypeName    string
	}{
		ContextArgName:           meta.ContextArg.Name,
		ErrorsRef:                meta.ErrorsArg.Name,
		RepoRef:                  meta.RepoArg.Name,
		IdParamName:              meta.ParamNameId,
		SingleName:               meta.SingleName,
		ResourceResponseTypeName: meta.ResourceResponseType.Name,
		UpdateRequestTypeName:    meta.RequestUpdateType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.ContextArg)

	handler.AddArgV(meta.ErrorsArg)
	handler.AddArgV(meta.RepoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		verb:          verbPut,
		uri:           fmt.Sprintf(`"/:%s"`, meta.ParamNameId),
		handlerFunc:   handler,
		requestStruct: nil,
	}
}
