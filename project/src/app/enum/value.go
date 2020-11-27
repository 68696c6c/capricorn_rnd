package enum

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

func makeValue(o config.EnumOptions, meta enumMeta) *golang.Function {
	method := golang.NewFunction(o.ValueFuncName)
	t := `
	return {{ .ReceiverName }}.String(), nil
`

	method.AddReturn("", golang.MakeTypeDriverValue())
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ReceiverName string
	}{
		ReceiverName: meta.enumType.GetReceiverName(),
	})

	return method
}
