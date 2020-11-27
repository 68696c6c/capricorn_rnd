package enum

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

func makeScan(o config.EnumOptions, meta enumMeta) *golang.Function {
	method := golang.NewFunction(o.ScanFuncName)
	t := `
	stringValue := fmt.Sprintf("%v", {{ .ArgName }})
	result, err := {{ .FromStringFuncName }}(stringValue)
	if err != nil {
		return err
	}
	*i = result
	return nil
`
	method.AddArg(o.InputArgName, golang.MakeTypeInterfaceLiteral())

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ArgName            string
		FromStringFuncName string
		BaseQueryFuncCall  string
	}{
		ArgName:            o.InputArgName,
		FromStringFuncName: meta.fromStringFuncName,
	})

	method.AddImportsStandard("fmt")

	return method
}
