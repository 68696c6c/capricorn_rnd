package golang

import "fmt"

type TypeAlias struct {
	*Type
	alias IType
}

func NewTypeAlias(name string, aliasType IType) *TypeAlias {
	return &TypeAlias{
		Type:  NewType(name, false, false),
		alias: aliasType,
	}
}

func (t TypeAlias) Render() string {
	aliasRef := t.alias.GetReference()
	if t.GetPackage() == t.alias.GetPackage() {
		aliasRef = t.alias.GetName()
	}
	return fmt.Sprintf("type %s %s", t.Name, aliasRef)
}
