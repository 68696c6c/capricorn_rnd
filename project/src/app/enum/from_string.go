package enum

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func makeFromString(o config.EnumOptions, meta enumMeta) *golang.Function {
	method := golang.NewFunction(meta.fromStringFuncName)
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
	method.AddArg(o.InputArgName, golang.MakeTypeString(false))

	method.AddReturn("", meta.enumType)
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ArgName          string
		ReadableTypeName string
		Values           []string
	}{
		ArgName:          o.InputArgName,
		ReadableTypeName: utils.Space(meta.name),
		Values:           meta.enumType.GetValues(),
	})

	method.AddImportsStandard("database/sql/driver", "fmt")
	method.AddImportsVendor(goat.ImportErrors)

	return method
}
