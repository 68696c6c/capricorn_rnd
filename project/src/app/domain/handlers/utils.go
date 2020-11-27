package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const (
	verbGet    = "GET"
	verbPost   = "POST"
	verbPut    = "PUT"
	verbDelete = "DELETE"
)

type handlerMeta struct {
	contextArg           *golang.Value
	errorsArg            *golang.Value
	repoArg              *golang.Value
	resourceName         string
	nameSingular         string
	namePlural           string
	modelTypeName        string
	requestCreateType    *golang.Struct
	requestUpdateType    *golang.Struct
	resourceResponseType *golang.Struct
	listResponseType     *golang.Struct
}

func makeHandlerMeta(o config.HandlersOptions, domainMeta *config.DomainMeta) handlerMeta {
	modelType := domainMeta.GetModelType()
	repoArgName := utils.Camel(o.RepoArgNameTemplate.Parse(domainMeta.ResourceName))

	createRequestName := utils.Pascal(o.CreateRequestNameTemplate.Parse(domainMeta.ResourceName))
	updateRequestName := utils.Pascal(o.UpdateRequestNameTemplate.Parse(domainMeta.ResourceName))

	resourceResponseName := utils.Camel(o.ResourceResponseNameTemplate.Parse(domainMeta.ResourceName))
	listResponseName := utils.Camel(o.ListResponseNameTemplate.Parse(domainMeta.ResourceName))

	domainRepo := domainMeta.GetRepoType()

	return handlerMeta{
		contextArg:           golang.ValueFromType(o.ContextArgName, goat.MakeTypeGinContext()),
		errorsArg:            golang.ValueFromType(o.ErrorsArgName, goat.MakeTypeErrorHandler()),
		repoArg:              golang.ValueFromType(repoArgName, domainRepo),
		resourceName:         domainMeta.ResourceName,
		nameSingular:         domainMeta.NameSingular,
		namePlural:           domainMeta.NamePlural,
		modelTypeName:        modelType.GetName(),
		requestCreateType:    makeCreateRequest(createRequestName, modelType),
		requestUpdateType:    makeCreateRequest(updateRequestName, modelType),
		resourceResponseType: makeResourceResponse(resourceResponseName, modelType),
		listResponseType:     makeListResponse(listResponseName, modelType),
	}
}

func makeHandlerFunc(handlerName, templateBody string, templateData interface{}, contextArg *golang.Value) *golang.Function {
	handler := golang.NewFunction(handlerName)
	t := `
	return {{ .InnerFunc.Render }}
`
	innerFunc := golang.NewFunction("")
	innerFunc.AddArgV(contextArg)
	innerFunc.SetBodyTemplate(templateBody, templateData)

	handler.AddReturn("", goat.MakeTypeHandlerFunc())

	handler.SetBodyTemplate(t, struct {
		InnerFunc utils.Renderable
	}{
		InnerFunc: innerFunc,
	})

	handler.AddImportsVendor(goat.ImportGoat, goat.ImportGin)

	return handler
}
