package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Handlers struct {
	*golang.File
	uri       string
	endpoints []*Handler
}

type Handler struct {
	*golang.Function
	verb          string
	uri           string
	requestStruct *golang.Struct
}

func (h *Handler) GetVerb() string {
	return h.verb
}

func (h *Handler) GetUri() string {
	return h.uri
}

func (h *Handler) HasRequest() bool {
	return h.requestStruct != nil
}

func (h *Handler) GetRequestStruct() golang.IType {
	return h.requestStruct
}

func Build(pkg golang.IPackage, fileName string, domainMeta *config.DomainMeta) *Handlers {
	actions := domainMeta.GetHandlerActions()
	if len(actions) == 0 {
		return nil
	}

	result := &Handlers{
		File:      pkg.AddGoFile(fileName),
		uri:       fmt.Sprintf("/%s", utils.Kebob(domainMeta.NamePlural)),
		endpoints: []*Handler{},
	}

	meta := makeHandlerMeta(domainMeta)

	var needResourceResponse bool
	for _, a := range actions {
		switch a {

		case config.ActionCreate:
			h := makeCreate(meta)
			result.AddStruct(meta.RequestCreateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionUpdate:
			h := makeUpdate(meta)
			result.AddStruct(meta.RequestUpdateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionView:
			h := makeView(meta)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionList:
			h := makeList(meta)
			result.AddStruct(meta.ListResponseType)
			result.AddFunction(h.Function)
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionDelete:
			h := makeDelete(meta)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break
		}
	}

	if needResourceResponse {
		result.AddStruct(meta.ResourceResponseType)
	}

	return result
}

func (g *Handlers) GetEndpoints() []*Handler {
	return g.endpoints
}

func (g *Handlers) GetUri() string {
	return g.uri
}
