package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

type Repo struct {
	*golang.File
	constructor    *golang.Function
	interfaceType  *golang.Type
	pageFuncName   string
	filterFuncName string
}

func NewRepo(pkg golang.IPackage, fileName string, domainMeta *config.DomainResource) *Repo {
	actions := domainMeta.GetRepoActions()
	if len(actions) == 0 {
		return nil
	}
	repoStruct, repoInterface, c := newRepoTypes(fileName, domainMeta)

	result := &Repo{
		File:           pkg.AddGoFile(fileName),
		pageFuncName:   c.pageQueryFuncName,
		filterFuncName: c.filterFuncName,
	}
	result.AddStruct(repoStruct)
	result.constructor = repoStruct.GetConstructor()

	result.AddInterface(repoInterface)
	result.interfaceType = repoInterface.CopyType()

	return result
}

func (r *Repo) GetInterfaceType() *golang.Type {
	return r.interfaceType.CopyType()
}

func (r *Repo) GetConstructor() *golang.Function {
	return r.constructor
}

func (r *Repo) GetPaginationFuncName() string {
	return r.pageFuncName
}

func (r *Repo) GetFilterFuncName() string {
	return r.filterFuncName
}
