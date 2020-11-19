package golang

import "fmt"

type Var struct {
	*Type
	Name     string
	Value    string
	implicit bool
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

func (v *Var) Render() string {
	hasValue := v.Value != ""
	if v.implicit && hasValue {
		return fmt.Sprintf("%s := %s", v.Name, v.Value)
	}
	if hasValue {
		return fmt.Sprintf("var %s = %s", v.Name, v.Value)
	}
	return fmt.Sprintf("var %s %s", v.Name, v.Type.GetReference())
}
