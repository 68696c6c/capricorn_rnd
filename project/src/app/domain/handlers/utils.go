package handlers

import (
	"fmt"

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

	paramNameId = "id"
)

type handlerMeta struct {
	ContextArg           *golang.Value
	ErrorsArg            *golang.Value
	RepoArg              *golang.Value
	SingleName           string
	PluralName           string
	ModelTypeName        string
	RequestCreateType    *golang.Struct
	RequestUpdateType    *golang.Struct
	ResourceResponseType *golang.Struct
	ListResponseType     *golang.Struct
	RepoPageFuncName     string
	RepoFilterFuncName   string
	ParamNameId          string
}

func makeHandlerMeta(domainMeta *config.DomainResource) handlerMeta {
	modelType := domainMeta.GetModelType()
	repoArgName := fmt.Sprintf("%sRepo", utils.Camel(domainMeta.NamePlural))
	return handlerMeta{
		ContextArg:           golang.ValueFromType("c", goat.MakeTypeGinContext()),
		ErrorsArg:            golang.ValueFromType("errorHandler", goat.MakeTypeErrorHandler()),
		RepoArg:              golang.ValueFromType(repoArgName, domainMeta.GetRepoType()),
		SingleName:           domainMeta.NameSingular,
		PluralName:           domainMeta.NamePlural,
		ModelTypeName:        modelType.GetName(),
		RequestCreateType:    makeCreateRequest("CreateRequest", modelType),
		RequestUpdateType:    makeCreateRequest("UpdateRequest", modelType),
		ResourceResponseType: makeResourceResponse("resourceResponse", modelType),
		ListResponseType:     makeListResponse("listResponse", modelType),
		RepoPageFuncName:     domainMeta.RepoPaginationFuncName,
		RepoFilterFuncName:   domainMeta.RepoFilterFuncName,
		ParamNameId:          paramNameId,
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
