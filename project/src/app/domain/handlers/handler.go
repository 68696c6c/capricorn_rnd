package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const (
	MiddlewareKeyAuth         = "auth"
	MiddlewareKeyBind         = "bind"
	MiddlewareKeyInitRegistry = "registry"
)

type Endpoints map[string]Handler

type Handlers struct {
	*golang.File
	endpoints Endpoints
}

type middlewares map[string]*golang.Function

type Handler struct {
	handlerFunc *golang.Function
	middlewares
}

type RouteGroup struct {
	Name        string
	Uri         string
	Middlewares middlewares
}

type handlerGroupMeta struct {
	ContextArg               *golang.Value
	ErrorsArg                *golang.Value
	RepoArg                  *golang.Value
	SingleName               string
	ModelType                *golang.Struct
	RequestCreateTypeName    string
	RequestUpdateTypeName    string
	ResourceResponseTypeName string
	ListResponseTypeName     string
}

func NewHandlers(pkg golang.IPackage, fileName string, meta model.Meta, repoType *golang.Interface) Handlers {
	result := Handlers{
		File: pkg.AddGoFile(fileName),
	}

	repoArgName := fmt.Sprintf("%sRepo", utils.Camel(meta.PluralName))
	handlerMeta := handlerGroupMeta{
		ContextArg:               golang.ValueFromType("c", goat.MakeTypeGinContext()),
		ErrorsArg:                golang.ValueFromType("errorHandler", goat.MakeTypeErrorHandler()),
		RepoArg:                  golang.ValueFromType(repoArgName, repoType.Type),
		SingleName:               meta.SingleName,
		ModelType:                meta.ModelType.Struct,
		RequestCreateTypeName:    fmt.Sprintf("CreateRequest"),
		RequestUpdateTypeName:    fmt.Sprintf("UpdateRequest"),
		ResourceResponseTypeName: fmt.Sprintf("resourceResponse"),
		ListResponseTypeName:     fmt.Sprintf("listResponse"),
	}

	var needResourceResponse bool
	for _, a := range meta.Actions {
		switch a {
		case model.ActionCreate:
			h := makeCreate(handlerMeta)
			result.AddStruct(makeCreateRequest(handlerMeta))
			result.AddFunction(h.handlerFunc)
			needResourceResponse = true
			break
			// case model.ActionUpdate:
			// 	m := makeUpdate(meta)
			// 	result.AddFunction(m)
			// 	break
			// case model.ActionView:
			// 	m := makeGetSingle(handlerMeta)
			// 	result.AddStruct(make)
			// 	result.AddFunction(m.handlerFunc)
			// 	break
			// case model.ActionList:
			// 	m := makeGetAll(meta)
			// 	result.AddFunction(m)
			// 	break
			// case model.ActionDelete:
			// 	m := makeDelete(meta)
			// 	result.AddFunction(m)
			// 	break
		}
	}

	if needResourceResponse {
		result.AddStruct(makeResourceResponse(handlerMeta))
	}

	return result
}
