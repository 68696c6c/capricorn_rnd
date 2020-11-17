package golang

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Functions []*Function

type Function struct {
	*Type
	arguments    []*Value
	returns      []*Value
	bodyTemplate string
	bodyData     interface{}
}

func NewFunction(name string) *Function {
	funcType := NewType(name, false, false)
	// The Type constructor sets a default receiver because that is usually helpful, but by default a function should not
	// have a receiver.
	funcType.SetReceiver(nil)
	return &Function{
		Type:         funcType,
		bodyTemplate: "",
	}
}

func (f *Function) GetType() *Type {
	return f.Type
}

func (f *Function) AddArg(name string, t IType) {
	f.arguments = append(f.arguments, ValueFromType(name, t.GetType()))
}

func (f *Function) AddReturn(name string, t IType) {
	f.returns = append(f.returns, ValueFromType(name, t.GetType()))
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
	if len(f.returns) > 1 || hasNamedReturns {
		returns = fmt.Sprintf("(%s)", returns)
	}
	result := fmt.Sprintf("%s(%s) %s", f.Name, args, returns)
	return strings.TrimSpace(result)
}

func (f *Function) getReceiver() string {
	if f.receiver == nil {
		return ""
	}
	r := fmt.Sprintf("%s %s", f.receiver.Name, f.receiver.GetReference())
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
		builtValues = append(builtValues, fmt.Sprintf("func %s%s {%s}\n", rec, sig, function.Render()))
	}
	if len(builtValues) == 0 {
		return ""
	}
	joinedValues := strings.Join(builtValues, "\n")
	result := strings.TrimSpace(joinedValues)
	return result
}
