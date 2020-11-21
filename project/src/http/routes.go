package http

import (
	"fmt"
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/handlers"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Routes struct {
	*golang.File
}

func buildRoutes(pkg *golang.Package, a app.App) Routes {
	result := Routes{
		File: pkg.AddGoFile("routes"),
	}

	initRouterFunc := golang.NewFunction("InitRouter")
	t := `
	router := goat.GetRouter()
	engine := router.GetEngine()

	engine.GET("/health", func(c *gin.Context) {
		goat.RespondMessage(c, "ok")
	})
	engine.GET("/version", func(c *gin.Context) {
		// TODO: show version.
		goat.RespondMessage(c, "something helpful here")
	})
	api := engine.Group("/api")

{{ .Groups.Render }}

	return router
`

	servicesArgName := "s"
	serviceContainerType := a.Container.GetContainerType()
	initRouterFunc.AddArg(servicesArgName, serviceContainerType)

	initRouterFunc.AddReturn("", goat.MakeTypeRouter())

	result.AddImportsApp(serviceContainerType.GetImport())

	errorsRef := fmt.Sprintf("%s.%s", servicesArgName, a.Container.ErrorHandlerField().Name)
	var groups handlers.RouteGroups
	for domainKey, d := range a.Domains {
		if !d.HasHandlers() {
			continue
		}
		d.Handlers.SetErrorsRef(errorsRef)

		repoField, err := a.Container.GetDomainRepoField(domainKey)
		if err != nil {
			panic(err)
		}

		d.Handlers.SetRepoRef(fmt.Sprintf("%s.%s", servicesArgName, repoField.Name))
		groups = append(groups, d.Handlers)
		result.AddImportsApp(d.GetImport())
	}

	initRouterFunc.SetBodyTemplate(t, struct {
		Groups utils.Renderable
	}{
		Groups: groups,
	})

	initRouterFunc.AddImportsVendor(goat.ImportGoat)

	result.AddFunction(initRouterFunc)

	result.AddImportsVendor(goat.ImportGin)

	return result
}
