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

func Build(pkg golang.IPackage, o config.HandlersOptions, domainMeta *config.DomainMeta) *Handlers {
	actions := domainMeta.GetHandlerActions()
	if len(actions) == 0 {
		return nil
	}

	fileName := o.FileNameTemplate.Parse(domainMeta.ResourceName)
	uri := o.UriTemplate.Parse(domainMeta.ResourceName)
	result := &Handlers{
		File:      pkg.AddGoFile(fileName),
		uri:       fmt.Sprintf("/%s", utils.Kebob(uri)),
		endpoints: []*Handler{},
	}

	meta := makeHandlerMeta(o, domainMeta)

	var needResourceResponse bool
	for _, a := range actions {
		switch a {

		case config.ActionCreate:
			h := makeCreate(o, meta)
			result.AddStruct(meta.requestCreateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionUpdate:
			h := makeUpdate(o, meta)
			result.AddStruct(meta.requestUpdateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionView:
			h := makeView(o, meta)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionList:
			h := makeList(o, meta)
			result.AddStruct(meta.listResponseType)
			result.AddFunction(h.Function)
			result.endpoints = append(result.endpoints, h)
			break

		case config.ActionDelete:
			h := makeDelete(o, meta)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break
		}
	}

	if needResourceResponse {
		result.AddStruct(meta.resourceResponseType)
	}

	result.AddImportsApp(domainMeta.ImportRepos, domainMeta.ImportModels)

	return result
}

func (g *Handlers) GetEndpoints() []*Handler {
	return g.endpoints
}

func (g *Handlers) GetUri() string {
	return g.uri
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
