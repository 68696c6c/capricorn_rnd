package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func newRepoTypes(fileName string, domainMeta *config.DomainResource) (*golang.Struct, *golang.Interface, methodMeta) {
	baseTypeName := utils.Pascal(fileName)
	repoStruct := golang.NewStruct(baseTypeName+"Gorm", false, false)
	repoInterface := golang.NewInterface(baseTypeName, false, false)

	meta := makeMethodMeta(domainMeta, repoStruct.GetReceiverName(), repoStruct.CopyType(), repoInterface.CopyType())

	repoStruct.AddField(golang.NewField(meta.dbFieldName, meta.dbType, false))

	repoStruct.AddConstructor(makeConstructor(meta))

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
		repoStruct.AddFunction(makeGetBaseQuery(meta))
		repoStruct.AddFunction(makeGetFilteredQuery(meta))

		m := makeApplyPaginationToQuery(meta)
		repoStruct.AddFunction(m)
		repoInterface.AddFunction(m)
	}

	return repoStruct, repoInterface, meta
}
