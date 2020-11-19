package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func newRepoTypes(fileName string, modelType model.Type, actions []model.Action) (*golang.Struct, *golang.Interface) {
	baseTypeName := utils.Pascal(fileName)
	repoStruct := golang.NewStruct(baseTypeName+"Gorm", false, false)
	repoInterface := golang.NewInterface(baseTypeName, false, false)

	meta := makeMethodMeta(modelType, repoStruct.GetReceiverName(), repoStruct.CopyType(), repoInterface.CopyType())

	repoStruct.AddField(golang.NewField(meta.dbFieldName, meta.dbType, false))

	repoStruct.AddConstructor(makeConstructor(meta))

	var needFilterFuncs bool
	var saveDone bool
	for _, a := range actions {
		switch a {
		case model.ActionCreate:
			fallthrough
		case model.ActionUpdate:
			if !saveDone {
				m := makeSave(meta)
				repoStruct.AddFunction(m)
				repoInterface.AddFunction(m)
			}
			saveDone = true
			break
		case model.ActionView:
			m := makeGetById(meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			break
		case model.ActionList:
			m := makeFilter(meta)
			repoStruct.AddFunction(m)
			repoInterface.AddFunction(m)
			needFilterFuncs = true
			break
		case model.ActionDelete:
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
		repoStruct.AddFunction(makeApplyPaginationToQuery(meta))
	}

	return repoStruct, repoInterface
}
