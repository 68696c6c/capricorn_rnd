package enum

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Map map[string]*golang.Iota

type Enums struct {
	*golang.Package
	enums   Map
	options config.EnumOptions
}

func NewEnums(pkg *golang.Package, o config.EnumOptions, enums []config.Enum) Enums {
	enumMap := make(Map)
	pkgEnums := pkg.AddPackage(o.PkgName)
	for _, e := range enums {
		key := utils.Snake(e.Name)
		enumMap[key] = newEnum(pkgEnums, o, e)
	}
	return Enums{
		Package: pkgEnums,
		enums:   enumMap,
		options: o,
	}
}

func newEnum(pkg *golang.Package, o config.EnumOptions, e config.Enum) *golang.Iota {
	fileName := utils.Snake(e.Name)
	typeName := utils.Pascal(e.Name)

	file := pkg.AddGoFile(fileName)

	enumType := golang.NewIota(typeName, e.Values)

	meta := enumMeta{
		name:               e.Name,
		fromStringFuncName: fmt.Sprintf("%s%s", typeName, o.FromStringFuncNameSuffix),
		enumType:           enumType,
	}

	file.AddFunction(makeFromString(o, meta))
	enumType.AddFunction(makeScan(o, meta))
	enumType.AddFunction(makeValue(o, meta))

	file.AddIota(enumType)

	return enumType
}

func (e Enums) GetEnumType(input string) (golang.IType, bool) {
	key := input
	if strings.HasPrefix(key, e.options.SpecPrefix) {
		key = strings.TrimPrefix(key, e.options.SpecPrefix)
	}
	result, ok := e.enums[key]
	if !ok {
		return nil, false
	}
	return result, true
}

type enumMeta struct {
	name               string
	fromStringFuncName string
	enumType           *golang.Iota
}
