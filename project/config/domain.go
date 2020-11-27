package config

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type DomainMeta struct {
	ModelType golang.IType
	RepoType  golang.IType

	ResourceName string
	NameSingular string
	NamePlural   string

	RepoActions    []Action
	HandlerActions []Action
	ServiceActions []string

	ImportModels   string
	ImportRepos    string
	ImportServices string
	ImportHandlers string
}

func NewDomainMeta(resourceName string, resourceActions []Action, customActions []string, pkgModels, pkgRepos, pkgServices, pkgHandlers golang.IPackage) *DomainMeta {
	repoActions, handlerActions := getResourceActions(resourceActions)
	return &DomainMeta{
		RepoType:       nil,
		ModelType:      nil,
		ResourceName:   resourceName,
		NameSingular:   utils.Singular(resourceName),
		NamePlural:     utils.Plural(resourceName),
		RepoActions:    repoActions,
		HandlerActions: handlerActions,
		ServiceActions: customActions,
		ImportModels:   pkgModels.GetImport(),
		ImportRepos:    pkgRepos.GetImport(),
		ImportServices: pkgServices.GetImport(),
		ImportHandlers: pkgHandlers.GetImport(),
	}
}

func (r *DomainMeta) SetModel(modelType golang.IType) {
	r.ModelType = modelType
}

func (r *DomainMeta) SetRepo(repoType golang.IType) {
	r.RepoType = repoType
}

func (r *DomainMeta) GetModelType() *golang.Type {
	return r.ModelType.CopyType()
}

func (r *DomainMeta) GetRepoType() *golang.Type {
	return r.RepoType.CopyType()
}

func (r *DomainMeta) GetRepoActions() []Action {
	return r.RepoActions
}

func (r *DomainMeta) GetHandlerActions() []Action {
	return r.HandlerActions
}

func (r *DomainMeta) GetServiceActions() []string {
	return r.ServiceActions
}

type DomainOptions struct {
	Model    ModelOptions
	Repo     RepoOptions
	Service  ServiceOptions
	Handlers HandlersOptions
}
