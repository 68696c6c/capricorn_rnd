package repo

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

type Repo struct {
	file *golang.File
}

func NewRepo(fileName string, meta model.Meta) Repo {
	return Repo{
		file: meta.PKG.AddGoFile(fileName),
	}
}
