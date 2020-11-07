package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type HTTP struct {
	*golang.Package
	Serve Serve
}

func Build(pkg *golang.Package, a app.App) {
	pkgHttp := pkg.AddPackage("http")
	buildServe(pkgHttp, a)
}
