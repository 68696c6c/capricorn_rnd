package golang

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Var struct {
	*Type
	Name          string
	Value         string
	implicit      bool
	valueTemplate string
	valueData     interface{}
}

func NewVar(name, value string, t IType, implicit bool) *Var {
	return &Var{
		Type:     t.GetType(),
		Name:     name,
		Value:    value,
		implicit: implicit,
	}
}

func (v *Var) GetType() *Type {
	return v.Type
}

func (v *Var) CopyType() *Type {
	return copyType(v.Type)
}

func (v *Var) SetValueTemplate(t string, data interface{}) {
	v.valueTemplate = t
	v.valueData = data
}

func (v *Var) Render() string {

	// Var declaration with assignment to a templated value.
	if v.valueTemplate != "" {
		val, err := utils.ParseTemplate("template_var", v.valueTemplate, v.valueData)
		if err != nil {
			panic(err)
		}
		if v.implicit {
			return fmt.Sprintf("%s := %s", v.Name, val)
		}
		return fmt.Sprintf("var %s = %s", v.Name, val)
	}

	// Var declaration with assignment to a simple value.
	if v.Value != "" {
		if v.implicit {
			return fmt.Sprintf("%s := %s", v.Name, v.Value)
		}
		return fmt.Sprintf("var %s = %s", v.Name, v.Value)
	}

	// Var declaration without any assignment.
	return fmt.Sprintf("var %s %s", v.Name, v.Type.GetReference())
}
