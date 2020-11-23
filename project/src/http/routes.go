package http

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
)

type Routes struct {
	*golang.File
	apiRoutePrefix string
	apiGroupName   string
}

func buildRoutes(pkg *golang.Package, a *app.App) Routes {
	result := Routes{
		File: pkg.AddGoFile("routes"),
	}

	initRouterFunc := makeInitRouter(a)

	result.AddFunction(initRouterFunc)

	result.AddImportsVendor(goat.ImportGin)

	return result
}
