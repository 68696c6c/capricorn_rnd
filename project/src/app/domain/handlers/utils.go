package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/repo"
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
	ModelType            *golang.Struct
	RequestCreateType    *golang.Struct
	RequestUpdateType    *golang.Struct
	ResourceResponseType *golang.Struct
	ListResponseType     *golang.Struct
	RepoPageFuncName     string
	RepoFilterFuncName   string
	ParamNameId          string
}

func makeHandlerMeta(modelMeta model.Meta, domainRepo *repo.Repo) handlerMeta {
	repoArgName := fmt.Sprintf("%sRepo", utils.Camel(modelMeta.PluralName))
	return handlerMeta{
		ContextArg:           golang.ValueFromType("c", goat.MakeTypeGinContext()),
		ErrorsArg:            golang.ValueFromType("errorHandler", goat.MakeTypeErrorHandler()),
		RepoArg:              golang.ValueFromType(repoArgName, domainRepo.GetInterfaceType()),
		SingleName:           modelMeta.SingleName,
		PluralName:           modelMeta.PluralName,
		ModelType:            modelMeta.ModelType.Struct,
		RequestCreateType:    makeCreateRequest("CreateRequest", modelMeta.ModelType.Struct),
		RequestUpdateType:    makeCreateRequest("UpdateRequest", modelMeta.ModelType.Struct),
		ResourceResponseType: makeResourceResponse("resourceResponse", modelMeta.ModelType.Struct),
		ListResponseType:     makeListResponse("listResponse", modelMeta.ModelType.Struct),
		RepoPageFuncName:     domainRepo.GetPaginationFuncName(),
		RepoFilterFuncName:   domainRepo.GetFilterFuncName(),
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
