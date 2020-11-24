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
	pageFuncName   string
	filterFuncName string
}

func Build(pkg golang.IPackage, fileName string, domainMeta *config.DomainMeta) *Repo {
	actions := domainMeta.GetRepoActions()
	if len(actions) == 0 {
		return nil
	}

	meta := buildTypeMeta(fileName, domainMeta)

	repoFile := pkg.AddGoFile(fileName)
	repoFile.AddStruct(meta.repoStructType)
	repoFile.AddInterface(meta.repoInterfaceType)

	return &Repo{
		File:           repoFile,
		constructor:    meta.constructor,
		interfaceType:  meta.repoInterfaceType,
		pageFuncName:   meta.pageQueryFuncName,
		filterFuncName: meta.filterFuncName,
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

func buildTypeMeta(fileName string, domainMeta *config.DomainMeta) *methodMeta {
	baseTypeName := utils.Pascal(fileName)
	repoStruct := golang.NewStruct(baseTypeName+"Gorm", false, false)
	repoInterface := golang.NewInterface(baseTypeName, false, false)

	meta := makeMethodMeta(domainMeta, repoStruct, repoInterface)

	repoStruct.AddField(golang.NewField(meta.dbFieldName, meta.dbType, false))

	constructor := makeConstructor(meta)
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
				m := makeSave(meta)
				repoStruct.AddFunction(m)
				repoInterface.AddFunction(m)
			}
			saveDone = true
			break

		case config.ActionView:
			m := makeGetById(meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			break

		case config.ActionList:
			m := makeFilter(meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			needFilterFuncs = true
			break

		case config.ActionDelete:
			m := makeDelete(meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			break
		}
	}

	// Add the unexported filter helper methods last so that they appear at the bottom of the file, keeping exported the
	// interface methods at the top.
	if needFilterFuncs {
		m := makeApplyPaginationToQuery(meta)
		repoStruct.AddFunction(m)
		repoInterface.AddFunction(m)

		repoStruct.AddFunction(makeGetBaseQuery(meta))
		repoStruct.AddFunction(makeGetFilteredQuery(meta))
	}

	return meta
}
