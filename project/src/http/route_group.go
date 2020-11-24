package http

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/handlers"
)

type routeGroups []*routeGroup

type routeGroup struct {
	*handlers.Handlers
	name           string
	errorsRef      string
	repoRef        string
	parentGroupRef string
}

func (g routeGroups) Render() string {
	var result []string
	for _, group := range g {
		result = append(result, group.Render())
	}
	return strings.Join(result, "\n")
}

func (g *routeGroup) Render() string {
	result := []string{
		"\t{",
		fmt.Sprintf(`		%s := %s.Group("%s")`, g.name, g.parentGroupRef, g.GetUri()),
	}
	for _, h := range g.GetEndpoints() {
		var handlerResult []string
		if h.HasRequest() {
			handlerResult = append(handlerResult, fmt.Sprintf("goat.BindMiddleware(%s{})", h.GetRequestStruct().GetReference()))
		}
		handlerCall := fmt.Sprintf("%s(%s, %s)", h.GetReference(), g.errorsRef, g.repoRef)
		handlerResult = append(handlerResult, handlerCall)
		handlerChain := strings.Join(handlerResult, ", ")
		result = append(result, fmt.Sprintf("\t\t%s.%s(%s, %s)", g.name, h.GetVerb(), h.GetUri(), handlerChain))
	}
	result = append(result, "\t}")
	return strings.Join(result, "\n")
}
