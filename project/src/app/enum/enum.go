package enum

import (
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const (
	pkgNameEnums = "enums"
	specPrefix   = "enum:"
)

type Map map[string]Enum

type Enums struct {
	*golang.Package
	enums Map
}

type Enum struct {
	*golang.File `yaml:"-"`
	Name         string       `yaml:"name,omitempty"`
	Description  string       `yaml:"description,omitempty"`
	Type         string       `yaml:"type,omitempty"`
	Values       []string     `yaml:"values,omitempty"`
	enumType     *golang.Iota `yaml:"-"`
}

func NewEnums(pkg *golang.Package, enums []Enum) Enums {
	enumMap := make(Map)
	pkgEnums := pkg.AddPackage(pkgNameEnums)
	for _, e := range enums {
		key := utils.Snake(e.Name)
		enumMap[key] = newEnum(pkgEnums, e)
	}
	return Enums{
		Package: pkgEnums,
		enums:   enumMap,
	}
}

func newEnum(pkg *golang.Package, e Enum) Enum {
	fileName := utils.Snake(e.Name)
	typeName := utils.Pascal(e.Name)
	result := Enum{
		File:        pkg.AddGoFile(fileName),
		Name:        typeName,
		Description: e.Description,
		Type:        e.Type,
		Values:      e.Values,
		enumType:    golang.NewIota(typeName, e.Values),
	}

	fromStringFuncName := typeName + "FromString"
	result.AddFunction(makeFromString(fromStringFuncName, result.enumType))
	result.enumType.AddFunction(makeScan(fromStringFuncName))
	result.enumType.AddFunction(makeValue(result.enumType.GetReceiverName()))

	result.File.AddIota(result.enumType)

	return result
}

func (e Enums) GetEnumType(input string) (golang.IType, bool) {
	key := input
	if strings.HasPrefix(key, specPrefix) {
		key = strings.TrimPrefix(key, specPrefix)
	}
	result, ok := e.enums[key]
	if !ok {
		return nil, false
	}
	return result.enumType, true
}
