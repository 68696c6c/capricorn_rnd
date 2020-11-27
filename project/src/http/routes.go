package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

func buildRoutes(pkg golang.IPackage, o config.RoutesOptions, a *app.App) *golang.Function {
	initRouterFunc := makeInitRouter(o, a)

	file := pkg.AddGoFile(o.FileName)
	file.AddFunction(initRouterFunc)
	file.AddImportsVendor(goat.ImportGin)

	return initRouterFunc
}
