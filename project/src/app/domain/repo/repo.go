package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

type Repo struct {
	*golang.File
}

func NewRepo(fileName string, meta model.Meta) Repo {
	result := &Repo{
		File: meta.PKG.AddGoFile(fileName),
	}
	repoStruct, repoInterface := newRepo(meta, result.PKG.GetBaseImport(), result.PKG.GetName(), fileName)
	result.AddStruct(repoStruct)
	result.AddInterface(repoInterface)
	return *result
}
