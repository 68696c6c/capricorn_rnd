package http

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type InitRouter struct {
	*golang.Function
}

func makeInitRouter(a *app.App) *golang.Function {
	result := golang.NewFunction("InitRouter")
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

	apiRoutesGroupName := "api"
	groupVarName := "g"
	servicesArgName := "s"
	serviceContainerType := a.GetContainerType()

	result.AddArg(servicesArgName, serviceContainerType)

	result.AddReturn("", goat.MakeTypeRouter())

	result.AddImportsApp(serviceContainerType.GetImport())

	errorsRef := fmt.Sprintf("%s.%s", servicesArgName, a.GetErrorHandlerFieldName())
	var groups routeGroups
	for domainKey, d := range a.GetDomains() {
		if !d.HasHandlers() {
			continue
		}

		repoFieldName, err := a.GetDomainRepoFieldName(domainKey)
		if err != nil {
			panic(err)
		}
		repoRef := fmt.Sprintf("%s.%s", servicesArgName, repoFieldName)

		groups = append(groups, &routeGroup{
			Handlers:       d.GetHandlers(),
			name:           groupVarName,
			errorsRef:      errorsRef,
			repoRef:        repoRef,
			parentGroupRef: apiRoutesGroupName,
		})
		result.AddImportsApp(d.GetImport())
	}

	result.SetBodyTemplate(t, struct {
		ApiRoutesGroupName string
		Groups             utils.Renderable
	}{
		ApiRoutesGroupName: apiRoutesGroupName,
		Groups:             groups,
	})

	result.AddImportsVendor(goat.ImportGoat)

	return result
}
