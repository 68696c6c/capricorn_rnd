package service

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/project/src/app/domain/model"
)

type Service struct {
	file *golang.File
}

func NewService(fileName string, meta model.Meta) Service {
	return Service{
		file: meta.PKG.AddGoFile(fileName),
	}
}
