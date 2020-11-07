package repo

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/project/src/app/domain/model"
)

type Repo struct {
	file *golang.File
}

func NewRepo(fileName string, meta model.Meta) Repo {
	return Repo{
		file: meta.PKG.AddGoFile(fileName),
	}
}
