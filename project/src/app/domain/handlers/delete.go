package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeDelete(o config.HandlersOptions, meta handlerMeta) *Handler {
	name := utils.Pascal(o.DeleteNameTemplate.Parse(meta.resourceName))
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

		err = {{ .RepoRef }}.{{ .RepoDeleteFuncName }}(m)
		if err != nil{
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to delete { .SingleName }}", goat.RespondServerError)
			return
		}

		goat.RespondValid(c)
	`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		RepoGetByIdFuncName      string
		RepoDeleteFuncName       string
		IdParamName              string
		SingleName               string
		ResourceResponseTypeName string
		UpdateRequestTypeName    string
	}{
		ContextArgName:           meta.contextArg.Name,
		ErrorsRef:                meta.errorsArg.Name,
		RepoRef:                  meta.repoArg.Name,
		RepoGetByIdFuncName:      o.RepoGetByIdFuncName,
		RepoDeleteFuncName:       o.RepoDeleteFuncName,
		IdParamName:              o.ParamNameId,
		SingleName:               meta.nameSingular,
		ResourceResponseTypeName: meta.resourceResponseType.Name,
		UpdateRequestTypeName:    meta.requestUpdateType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.contextArg)

	handler.AddArgV(meta.errorsArg)
	handler.AddArgV(meta.repoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbDelete,
		uri:           fmt.Sprintf(`"/:%s"`, o.ParamNameId),
		requestStruct: nil,
	}
}
