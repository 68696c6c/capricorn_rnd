package enum

import "github.com/68696c6c/capricorn_rnd/golang"

func makeValue(receiverName string) *golang.Function {
	method := golang.NewFunction("Value")
	t := `
	return {{ .ReceiverName }}.String(), nil
`

	method.AddReturn("", golang.MakeTypeDriverValue())
	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ReceiverName string
	}{
		ReceiverName: receiverName,
	})

	return method
}
