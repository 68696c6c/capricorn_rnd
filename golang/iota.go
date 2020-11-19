package golang

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Iota struct {
	*TypeAlias
	values []string
}

func NewIota(name string, values []string) *Iota {
	typeName := utils.Pascal(name)

	typeAlias := NewTypeAlias(typeName, MakeTypeInt(false), false)

	// The String() method that stringer generates uses 'i' as the receiver name.
	// The Scan and Value methods are the only functions we generate and both require a pointer receiver.
	rec := NewTypeAlias(typeName, MakeTypeInt(false), true)
	rec.SetPackage("")
	typeAlias.SetReceiver(ValueFromType("i", rec.GetType()))

	var iotaValues []string
	for _, v := range values {
		iotaValues = append(iotaValues, fmt.Sprintf("%s%s", typeName, utils.Pascal(v)))
	}

	return &Iota{
		TypeAlias: typeAlias,
		values:    iotaValues,
	}
}

func (i *Iota) GetType() *Type {
	return i.Type
}

func (i *Iota) CopyType() *Type {
	return copyType(i.Type)
}

func (i *Iota) GetValues() []string {
	return i.values
}

func (i *Iota) Render() string {
	var valueLines []string

	if len(i.values) > 0 {
		valueLines = append(valueLines, "")
		valueLines = append(valueLines, "const (")
		for index, iotaValue := range i.values {
			line := "\t" + iotaValue
			if index == 0 {
				line = fmt.Sprintf("%s %s = iota + 1", line, i.Name)
			}
			valueLines = append(valueLines, line)
		}
		valueLines = append(valueLines, ")")
	}
	result := []string{
		"",
		fmt.Sprintf("//go:generate stringer -type=%s -trimprefix=%s", i.Name, i.Name),
		i.TypeAlias.Render(),
	}
	result = append(result, valueLines...)
	result = append(result, "")
	result = append(result, i.functions.Render())
	return strings.Join(result, "\n")
}
