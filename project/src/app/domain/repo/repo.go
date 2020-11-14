package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

type Repo struct {
	*golang.File
	interfaceType *golang.Interface
}

func NewRepo(pkg *golang.Package, fileName string, meta model.Meta) Repo {
	modelType := determineModelType(pkg.GetName(), meta.ModelType)
	repoStruct, repoInterface := newRepoTypes(fileName, modelType, meta.Actions)

	result := &Repo{
		File:          pkg.AddGoFile(fileName),
		interfaceType: repoInterface,
	}
	result.AddStruct(repoStruct)
	result.AddInterface(repoInterface)

	return *result
}

func (r Repo) GetInterfaceType() *golang.Interface {
	return r.interfaceType
}
