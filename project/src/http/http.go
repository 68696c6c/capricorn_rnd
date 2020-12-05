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

func Build(rootPath string, module string, o config.HttpOptions, a *app.App) *Http {
	pkgHttp := golang.NewPackage(o.PkgName, rootPath, module)
	return &Http{
		Package:    pkgHttp,
		initRouter: buildRoutes(pkgHttp, o.Routes, a),
	}
}

func (h *Http) GetInitRouter() *golang.Function {
	return h.initRouter
}
