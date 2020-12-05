package http

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeInitRouter(o config.RoutesOptions, a *app.App) *golang.Function {
	result := golang.NewFunction(o.RouterInitFuncName)
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
	{{ .ApiRoutesGroupName }} := engine.Group("/api")

{{ .Groups.Render }}

	return router
`
	appDomains := a.GetDomains()
	serviceContainerType := a.GetContainerType()

	result.AddArg(o.ServicesArgName, serviceContainerType)

	result.AddReturn("", goat.MakeTypeRouter())

	result.AddImportsApp(serviceContainerType.GetImport(), appDomains.GetImportHandlers())

	errorsRef := fmt.Sprintf("%s.%s", o.ServicesArgName, a.GetErrorHandlerFieldName())
	var groups routeGroups
	for domainKey, d := range appDomains.GetMap() {
		if !d.HasHandlers() {
			continue
		}

		repoFieldName, err := a.GetDomainRepoFieldName(domainKey)
		if err != nil {
			panic(err)
		}
		repoRef := fmt.Sprintf("%s.%s", o.ServicesArgName, repoFieldName)

		groups = append(groups, &routeGroup{
			Handlers:       d.GetHandlers(),
			name:           o.GroupVarName,
			errorsRef:      errorsRef,
			repoRef:        repoRef,
			parentGroupRef: o.ApiGroupName,
		})
	}

	result.SetBodyTemplate(t, struct {
		ApiRoutesGroupName string
		Groups             utils.Renderable
	}{
		ApiRoutesGroupName: o.ApiGroupName,
		Groups:             groups,
	})

	result.AddImportsVendor(goat.ImportGoat)

	return result
}
