package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type Http struct {
	*golang.Package
	initRouter *golang.Function
}

func Build(pkg golang.IPackage, o config.HttpOptions, a *app.App) *Http {
	pkgHttp := pkg.AddPackage(o.PkgName)
	return &Http{
		Package:    pkgHttp,
		initRouter: buildRoutes(pkgHttp, o.Routes, a),
	}
}

func (h *Http) GetInitRouter() *golang.Function {
	return h.initRouter
}
