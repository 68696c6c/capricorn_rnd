package enum

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeFromString(funcName string, enumType *golang.Iota) *golang.Function {
	argName := "input"
	method := golang.NewFunction(funcName)
	t := `
	switch {{ .ArgName }} {
	{{- range $key, $value := .Values }}
	case {{ $value }}.String():
		return {{ $value }}, nil
	{{- end }}
	default:
		return 0, errors.Errorf("invalid {{ .ReadableTypeName }} value '%s'", {{ .ArgName }})
	}
`

	method.AddArg(argName, golang.MakeTypeString(false))

	method.AddReturn("", enumType)
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ArgName          string
		ReadableTypeName string
		Values           []string
	}{
		ArgName:          argName,
		ReadableTypeName: utils.Space(enumType.GetName()),
		Values:           enumType.GetValues(),
	})

	method.AddImportsVendor(goat.ImportErrors)

	return method
}
