package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type Serve struct {
	*golang.File
}

func buildServe(pkg *golang.Package, a app.App) Serve {
	return Serve{
		File: pkg.AddGoFile("serve"),
	}
}
