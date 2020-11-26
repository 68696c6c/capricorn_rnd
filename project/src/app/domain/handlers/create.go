package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeCreateRequest(name string, modelType golang.IType) *golang.Struct {
	result := golang.NewStruct(name, false, false)
	result.AddField(golang.NewField("", modelType, true))
	return result
}

func makeCreate(meta handlerMeta) *Handler {
	name := fmt.Sprintf("Create%s", utils.Pascal(meta.SingleName))
	body := `
		req, ok := goat.GetRequest({{ .ContextArgName }}).(*{{ .RequestCreateTypeName }})
		if !ok {
			{{ .ErrorsRef }}.HandleMessage({{ .ContextArgName }}, "failed to get request", goat.RespondBadRequestError)
			return
		}

		m := req.{{ .ModelTypeName }}
		err := {{ .RepoRef }}.Save(&m)
		if err != nil {
			{{ .ErrorsRef }}.HandleErrorM({{ .ContextArgName }}, err, "failed to save {{ .SingleName }}", goat.RespondServerError)
			return
		}

		goat.RespondCreated({{ .ContextArgName }}, {{ .ResourceResponseTypeName }}{m})
	`
	data := struct {
		ContextArgName           string
		ErrorsRef                string
		RepoRef                  string
		SingleName               string
		ModelTypeName            string
		RequestCreateTypeName    string
		ResourceResponseTypeName string
	}{
		ContextArgName:           meta.ContextArg.Name,
		ErrorsRef:                meta.ErrorsArg.Name,
		RepoRef:                  meta.RepoArg.Name,
		SingleName:               utils.Space(meta.SingleName),
		ModelTypeName:            meta.ModelTypeName,
		RequestCreateTypeName:    meta.RequestCreateType.Name,
		ResourceResponseTypeName: meta.ResourceResponseType.Name,
	}

	handler := makeHandlerFunc(name, body, data, meta.ContextArg)

	handler.AddArgV(meta.ErrorsArg)
	handler.AddArgV(meta.RepoArg)

	handler.AddImportsVendor(goat.ImportGoat)

	return &Handler{
		Function:      handler,
		verb:          verbPost,
		uri:           `""`,
		requestStruct: meta.RequestCreateType,
	}
}
