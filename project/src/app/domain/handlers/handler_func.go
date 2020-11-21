package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeHandlerFunc(handlerName, templateBody string, templateData interface{}, contextArg *golang.Value) *golang.Function {
	handler := golang.NewFunction(handlerName)
	t := `
	return {{ .InnerFunc.Render }}
`
	innerFunc := golang.NewFunction("")
	innerFunc.AddArgV(contextArg)
	innerFunc.SetBodyTemplate(templateBody, templateData)

	handler.AddReturn("", goat.MakeTypeHandlerFunc())

	handler.SetBodyTemplate(t, struct {
		InnerFunc utils.Renderable
	}{
		InnerFunc: innerFunc,
	})

	handler.AddImportsVendor(goat.ImportGoat, goat.ImportGin)

	return handler
}
