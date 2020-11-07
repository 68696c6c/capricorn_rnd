package http

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/project/src/app"
)

type Serve struct {
	file *golang.File
}

func NewServe(pkg *golang.Package, a app.App) Serve {
	return Serve{
		file: pkg.AddGoFile("serve"),
	}
}
