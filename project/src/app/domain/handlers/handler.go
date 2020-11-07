package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
)

type Endpoints map[string]Handler

type Handlers struct {
	file      *golang.File
	endpoints Endpoints
}

type middlewares map[string]handlerFunc

type Handler struct {
	handlerFunc
	middlewares
}

type handlerFunc struct {
	Args []string
}

func NewHandlers(fileName string, meta model.Meta) Handlers {
	return Handlers{
		file: meta.PKG.AddGoFile(fileName),
	}
}
