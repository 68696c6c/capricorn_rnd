package repo

import (
	"path"

	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const repoReceiverName = "r"

type repoStruct struct {
	*golang.Struct
	baseImport string
	pkgName    string
	modelType  model.Type

	// Methods
	filter            *golang.Function
	getById           *golang.Function
	save              *golang.Function
	delete            *golang.Function
	getBaseQuery      *golang.Function
	getFilteredQuery  *golang.Function
	getPaginatedCount *golang.Function
}

func newRepo(baseImport, pkgName, fileName string, meta model.Meta) *repoStruct {
	result := &repoStruct{
		Struct: golang.NewStructFromType(golang.Type{
			Import:    path.Join(baseImport, pkgName),
			Package:   pkgName,
			Name:      utils.Pascal(fileName),
			IsPointer: false,
			IsSlice:   false,
		}),
		baseImport: baseImport,
		pkgName:    pkgName,
		modelType:  meta.ModelType,
	}
	result.SetReceiverName(repoReceiverName)
	result.SetReceiverTypeRef(result.Name)

	for _, a := range meta.Actions {
		switch a {
		case model.ActionList:
			result.AddFilterFunc()
			break
		}
	}

	return result
}

func (r *repoStruct) AddFilterFunc() {
	filter := makeFilter(r.baseImport, r.pkgName, r.GetReceiver().Name, r.modelType)
	r.AddFunction(filter)
}
