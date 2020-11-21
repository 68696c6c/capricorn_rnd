package handlers

import (
	"fmt"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/repo"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const (
	verbGet     = "GET"
	verbPost    = "POST"
	verbPut     = "PUT"
	verbDelete  = "DELETE"
	paramNameId = "id"
)

type RouteGroups []*RouteGroup

type RouteGroup struct {
	*golang.File
	name      string
	uri       string
	endpoints []*Handler
	errorsRef string
	repoRef   string
}

type Handler struct {
	verb          string
	uri           string
	handlerFunc   *golang.Function
	requestStruct *golang.Struct
}

func (h *Handler) renderHandlerChain(errorsRef, repoRef string) string {
	var result []string
	if h.requestStruct != nil {
		result = append(result, fmt.Sprintf("goat.BindMiddleware(%s{})", h.requestStruct.GetReference()))
	}
	handlerCall := fmt.Sprintf("%s(%s, %s)", h.handlerFunc.GetReference(), errorsRef, repoRef)
	result = append(result, handlerCall)
	return strings.Join(result, ", ")
}

func (g *RouteGroup) SetErrorsRef(ref string) {
	g.errorsRef = ref
}

func (g *RouteGroup) SetRepoRef(ref string) {
	g.repoRef = ref
}

func (g *RouteGroup) Render() string {
	result := []string{
		"\t{",
		fmt.Sprintf(`		%s := api.Group("%s")`, g.name, g.uri),
	}
	for _, h := range g.endpoints {
		handlerChain := h.renderHandlerChain(g.errorsRef, g.repoRef)
		result = append(result, fmt.Sprintf("\t\t%s.%s(%s, %s)", g.name, h.verb, h.uri, handlerChain))
	}
	result = append(result, "\t}")
	return strings.Join(result, "\n")
}

func (g RouteGroups) Render() string {
	var result []string
	for _, group := range g {
		result = append(result, group.Render())
	}
	return strings.Join(result, "\n")
}

type handlerGroupMeta struct {
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

func NewRouteGroup(pkg golang.IPackage, fileName string, meta model.Meta, domainRepo *repo.Repo) *RouteGroup {
	name := "g"
	if meta.ModelType.Struct.Name != "" {
		name = strings.ToLower(meta.ModelType.Struct.Name[0:1])
	}
	result := &RouteGroup{
		File: pkg.AddGoFile(fileName),
		name: name,
		uri:  fmt.Sprintf("/%s", utils.Kebob(meta.PluralName)),
	}

	createRequest := makeCreateRequest("CreateRequest", meta.ModelType.Struct)
	updateRequest := makeCreateRequest("UpdateRequest", meta.ModelType.Struct)
	resourceResponse := makeResourceResponse("resourceResponse", meta.ModelType.Struct)
	listResponse := makeListResponse("listResponse", meta.ModelType.Struct)

	repoArgName := fmt.Sprintf("%sRepo", utils.Camel(meta.PluralName))
	handlerMeta := handlerGroupMeta{
		ContextArg:           golang.ValueFromType("c", goat.MakeTypeGinContext()),
		ErrorsArg:            golang.ValueFromType("errorHandler", goat.MakeTypeErrorHandler()),
		RepoArg:              golang.ValueFromType(repoArgName, domainRepo.GetInterfaceType()),
		SingleName:           meta.SingleName,
		PluralName:           meta.PluralName,
		ModelType:            meta.ModelType.Struct,
		RequestCreateType:    createRequest,
		RequestUpdateType:    updateRequest,
		ResourceResponseType: resourceResponse,
		ListResponseType:     listResponse,
		RepoPageFuncName:     domainRepo.GetPaginationFuncName(),
		RepoFilterFuncName:   domainRepo.GetFilterFuncName(),
		ParamNameId:          paramNameId,
	}

	var endpoints []*Handler
	var needResourceResponse bool
	var needListResponse bool
	for _, a := range meta.Actions {
		switch a {
		case model.ActionCreate:
			h := makeCreate(handlerMeta)
			result.AddStruct(createRequest)
			result.AddFunction(h.handlerFunc)
			needResourceResponse = true
			endpoints = append(endpoints, h)
			break
		case model.ActionUpdate:
			h := makeUpdate(handlerMeta)
			result.AddStruct(updateRequest)
			result.AddFunction(h.handlerFunc)
			needResourceResponse = true
			endpoints = append(endpoints, h)
			break
		case model.ActionView:
			h := makeView(handlerMeta)
			result.AddFunction(h.handlerFunc)
			needResourceResponse = true
			endpoints = append(endpoints, h)
			break
		case model.ActionList:
			h := makeList(handlerMeta)
			result.AddStruct(createRequest)
			result.AddFunction(h.handlerFunc)
			needListResponse = true
			endpoints = append(endpoints, h)
			break
		case model.ActionDelete:
			h := makeDelete(handlerMeta)
			result.AddStruct(updateRequest)
			result.AddFunction(h.handlerFunc)
			needResourceResponse = true
			endpoints = append(endpoints, h)
			break
		}
	}

	if needResourceResponse {
		result.AddStruct(resourceResponse)
	}

	if needListResponse {
		result.AddStruct(listResponse)
	}

	result.endpoints = endpoints
	return result
}
