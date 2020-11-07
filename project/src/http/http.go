package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type HTTP struct {
	pkg   *golang.Package
	Serve Serve
}

func NewHTTP(pkg *golang.Package, a app.App) HTTP {
	pkgHttp := pkg.AddPackage("http")
	return HTTP{
		pkg:   pkgHttp,
		Serve: NewServe(pkgHttp, a),
	}
}
