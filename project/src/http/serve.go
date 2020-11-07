package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type Serve struct {
	file *golang.File
}

func NewServe(pkg *golang.Package, a app.App) Serve {
	return Serve{
		file: pkg.AddGoFile("serve"),
	}
}
