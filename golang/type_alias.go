package golang

import "fmt"

type TypeAlias struct {
	*Type
	alias IType
}

func NewTypeAlias(name string, aliasType IType, isPointer bool) *TypeAlias {
	return &TypeAlias{
		Type:  NewType(name, isPointer, false),
		alias: aliasType,
	}
}

func (t *TypeAlias) GetType() *Type {
	return t.Type
}

func (t *TypeAlias) Render() string {
	aliasRef := t.alias.GetReference()
	if t.GetPackage() == t.alias.GetPackage() {
		aliasRef = t.alias.GetName()
	}
	return fmt.Sprintf("type %s %s", t.Name, aliasRef)
}
