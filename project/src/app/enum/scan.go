package enum

import "github.com/68696c6c/capricorn_rnd/golang"

func makeScan(fromStringFuncName string) *golang.Function {
	argName := "input"
	method := golang.NewFunction("Scan")
	t := `
	stringValue := fmt.Sprintf("%v", {{ .ArgName }})
	result, err := {{ .FromStringFuncName }}(stringValue)
	if err != nil {
		return err
	}
	*i = result
	return nil
`

	method.AddArg(argName, golang.MakeTypeInterfaceLiteral())

	method.AddReturn("", golang.MakeTypeError())

	method.SetBodyTemplate(t, struct {
		ArgName            string
		FromStringFuncName string
		BaseQueryFuncCall  string
	}{
		ArgName:            argName,
		FromStringFuncName: fromStringFuncName,
	})

	return method
}
