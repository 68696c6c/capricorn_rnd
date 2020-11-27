package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Repo struct {
	*golang.File
	constructor    *golang.Function
	interfaceType  golang.IType
	externalName   string
	pageFuncName   string
	filterFuncName string
}

func Build(pkg golang.IPackage, o config.RepoOptions, domainMeta *config.DomainMeta) *Repo {
	actions := domainMeta.GetRepoActions()
	if len(actions) == 0 {
		return nil
	}

	meta := buildTypeMeta(o, domainMeta)

	fileName := o.FileNameTemplate.Parse(domainMeta.ResourceName)
	repoFile := pkg.AddGoFile(fileName)
	repoFile.AddStruct(meta.repoStructType)
	repoFile.AddInterface(meta.repoInterfaceType)
	repoFile.AddImportsApp(domainMeta.ImportModels)

	return &Repo{
		File:           repoFile,
		constructor:    meta.constructor,
		interfaceType:  meta.repoInterfaceType,
		externalName:   o.ExternalNameTemplate.Parse(domainMeta.ResourceName),
		pageFuncName:   o.PaginationFuncName,
		filterFuncName: o.FilterFuncName,
	}
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

func (r *Repo) GetExternalName() string {
	return r.externalName
}

func buildTypeMeta(o config.RepoOptions, domainMeta *config.DomainMeta) *methodMeta {
	intTypeName := utils.Pascal(o.InterfaceNameTemplate.Parse(domainMeta.ResourceName))
	repoInterface := golang.NewInterface(intTypeName, false, false)

	impTypeName := utils.Pascal(o.ImplementationNameTemplate.Parse(domainMeta.ResourceName))
	repoStruct := golang.NewStruct(impTypeName, false, false)

	meta := makeMethodMeta(o, domainMeta, repoStruct, repoInterface)

	repoStruct.AddField(golang.NewField(o.DbFieldName, meta.dbType, false))

	constructor := makeConstructor(o, meta)
	repoStruct.AddConstructor(constructor)
	meta.AddConstructor(constructor)

	var needFilterFuncs bool
	var saveDone bool
	for _, a := range domainMeta.GetRepoActions() {
		switch a {

		case config.ActionCreate:
			fallthrough
		case config.ActionUpdate:
			if !saveDone {
				m := makeSave(o, meta)
				repoStruct.AddFunction(m)
				repoInterface.AddFunction(m)
			}
			saveDone = true
			break

		case config.ActionView:
			m := makeGetById(o, meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			break

		case config.ActionList:
			m := makeFilter(o, meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			needFilterFuncs = true
			break

		case config.ActionDelete:
			m := makeDelete(o, meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			break
		}
	}

	// Add the unexported filter helper methods last so that they appear at the bottom of the file, keeping exported the
	// interface methods at the top.
	if needFilterFuncs {
		m := makeApplyPaginationToQuery(o, meta)
		repoStruct.AddFunction(m)
		repoInterface.AddFunction(m)

		repoStruct.AddFunction(makeGetBaseQuery(o, meta))
		repoStruct.AddFunction(makeGetFilteredQuery(o, meta))
	}

	return meta
}
