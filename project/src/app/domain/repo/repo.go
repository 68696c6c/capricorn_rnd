package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

type Repo struct {
	*golang.File
	structType    *golang.Struct
	interfaceType *golang.Interface
}

func NewRepo(pkg golang.IPackage, fileName string, meta model.Meta) *Repo {
	repoStruct, repoInterface := newRepoTypes(fileName, meta.ModelType, meta.Actions)

	result := &Repo{
		File:          pkg.AddGoFile(fileName),
		structType:    repoStruct,
		interfaceType: repoInterface,
	}
	result.AddStruct(repoStruct)
	result.AddInterface(repoInterface)

	return result
}

func (r *Repo) GetInterfaceType() *golang.Interface {
	return r.interfaceType
}

func (r *Repo) GetConstructor() *golang.Function {
	return r.structType.GetConstructor()
}
