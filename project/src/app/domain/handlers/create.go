package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeCreateRequest(name string, modelType golang.IType) *golang.Struct {
	result := golang.NewStruct(name, false, false)
	result.AddField(golang.NewField("", modelType, true))
	return result
}

func makeCreate(o config.HandlersOptions, meta handlerMeta) *Handler {
	name := utils.Pascal(o.CreateNameTemplate.Parse(meta.resourceName))
	body := `
		req, ok := goat.GetRequest({{ .ContextArgName }}).(*{{ .RequestCreateTypeName }})
		if !ok {
			{{ .ErrorsRef }}.HandleMessage({{ .ContextArgName }}, "failed to get request", goat.RespondBadRequestError)
			return
		}

		err := {{ .RepoRef }}.{{ .RepoSaveFuncName }}(&req.{{ .ModelTypeName }})
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to save {{ .SingleName }}", goat.RespondServerError)
			return
		}

		goat.RespondCreated({{ .ContextArgName }}, {{ .ResourceResponseTypeName }}{req.{{ .ModelTypeName }}})
	`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		RepoSaveFuncName         string
		SingleName               string
		ModelTypeName            string
		RequestCreateTypeName    string
		ResourceResponseTypeName string
	}{
		ContextArgName:           meta.contextArg.Name,
		ErrorsRef:                meta.errorsArg.Name,
		RepoRef:                  meta.repoArg.Name,
		RepoSaveFuncName:         o.RepoSaveFuncName,
		SingleName:               utils.Space(meta.nameSingular),
		ModelTypeName:            meta.modelTypeName,
		RequestCreateTypeName:    meta.requestCreateType.Name,
		ResourceResponseTypeName: meta.resourceResponseType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.contextArg)

	handler.AddArgV(meta.errorsArg)
	handler.AddArgV(meta.repoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbPost,
		uri:           `""`,
		requestStruct: meta.requestCreateType,
	}
}
