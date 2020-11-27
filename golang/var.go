package golang

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Var struct {
	*Type
	value         string
	valueType     *Type
	implicit      bool
	valueTemplate string
	valueData     interface{}
}

func NewVar(name, value string, valueType IType, implicit bool) *Var {
	return &Var{
		Type:      NewType(name, false, false),
		value:     value,
		valueType: valueType.GetType(),
		implicit:  implicit,
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
			return fmt.Sprintf("\n%s := %s", v.Name, val)
		}
		return fmt.Sprintf("\nvar %s = %s", v.Name, val)
	}

	// Var declaration with assignment to a simple value.
	if v.value != "" {
		if v.implicit {
			return fmt.Sprintf("\n%s := %s", v.Name, v.value)
		}
		return fmt.Sprintf("\nvar %s = %s", v.Name, v.value)
	}

	// Var declaration without any assignment.
	return fmt.Sprintf("\nvar %s %s", v.Name, v.valueType.GetReference())
}
