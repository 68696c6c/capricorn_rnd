package golang

import (
	"fmt"
	"path"
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Functions []*Function

type Function struct {
	Type
	*imports
	arguments    []Value
	returns      []Value
	receiver     *Value
	bodyTemplate string
	bodyData     interface{}
}

func NewFunction(baseImport, pkgName, name string) *Function {
	return &Function{
		Type: Type{
			Import:    path.Join(baseImport, pkgName),
			Package:   pkgName,
			Name:      name,
			IsPointer: false,
			IsSlice:   false,
		},
		imports:      newImports(),
		bodyTemplate: "",
	}
}

func (f *Function) GetImports() imports {
	return *f.imports
}

func (f *Function) AddArg(name string, t IType) {
	f.arguments = append(f.arguments, Value{
		TypeRef: t.GetReference(),
		Name:    name,
	})
}

func (f *Function) AddReturn(name string, t IType) {
	f.returns = append(f.returns, Value{
		TypeRef: t.GetReference(),
		Name:    name,
	})
}

func (f *Function) SetReceiver(v Value) {
	f.receiver = &v
}

func (f *Function) SetBodyTemplate(t string, data interface{}) {
	f.bodyTemplate = t
	f.bodyData = data
}

func (f *Function) GetSignature() string {
	args := getJoinedValueString(f.arguments)
	returns := getJoinedValueString(f.returns)
	var hasNamedReturns bool
	for _, r := range f.returns {
		if r.Name != "" {
			hasNamedReturns = true
			break
		}
	}
	if len(f.returns) > 0 || hasNamedReturns {
		returns = fmt.Sprintf("(%s)", returns)
	}
	result := fmt.Sprintf("%s(%s) %s", f.Name, args, returns)
	return strings.TrimSpace(result)
}

func (f *Function) getReceiver() string {
	r := fmt.Sprintf("%s %s", f.receiver.Name, f.receiver.TypeRef)
	r = strings.TrimSpace(r)
	if r != "" {
		return fmt.Sprintf("(%s) ", r)
	}
	return ""
}

func (f *Function) Render() string {
	result, err := utils.ParseTemplate("template_function", f.bodyTemplate, f.bodyData)
	if err != nil {
		panic(err)
	}
	return result
}

func (f Functions) Render() string {
	var builtValues []string
	for _, function := range f {
		rec := function.getReceiver()
		sig := function.GetSignature()
		builtValues = append(builtValues, fmt.Sprintf("func %s%s {%s}", rec, sig, function.Render()))
	}
	if len(builtValues) == 0 {
		return ""
	}
	joinedValues := strings.Join(builtValues, "\n")
	result := strings.TrimSpace(joinedValues)
	return result
}
