package handlers

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/repo"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Group struct {
	*golang.File
	uri       string
	endpoints []*Handler
}

func NewGroup(pkg golang.IPackage, fileName string, modelMeta model.Meta, domainRepo *repo.Repo) *Group {
	result := &Group{
		File:      pkg.AddGoFile(fileName),
		uri:       fmt.Sprintf("/%s", utils.Kebob(modelMeta.PluralName)),
		endpoints: []*Handler{},
	}

	meta := makeHandlerMeta(modelMeta, domainRepo)

	var needResourceResponse bool
	for _, a := range modelMeta.Actions {
		switch a {

		case model.ActionCreate:
			h := makeCreate(meta)
			result.AddStruct(meta.RequestCreateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case model.ActionUpdate:
			h := makeUpdate(meta)
			result.AddStruct(meta.RequestUpdateType)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case model.ActionView:
			h := makeView(meta)
			result.AddFunction(h.Function)
			needResourceResponse = true
			result.endpoints = append(result.endpoints, h)
			break

		case model.ActionList:
			h := makeList(meta)
			result.AddStruct(meta.ListResponseType)
			result.AddFunction(h.Function)
			result.endpoints = append(result.endpoints, h)
			break

		case model.ActionDelete:
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

func (g *Group) GetEndpoints() []*Handler {
	return g.endpoints
}

func (g *Group) GetUri() string {
	return g.uri
}
