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
	*golang.File
	Name        string   `yaml:"name,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Values      []string `yaml:"values,omitempty"`
}

func NewEnums(pkg *golang.Package, enums []Enum) Enums {
	result := make(Map)
	pkgEnums := pkg.AddPackage(pkgNameEnums)
	for _, e := range enums {
		key := utils.Snake(e.Name)
		result[key] = newEnum(pkgEnums, e)
	}
	return Enums{
		Package: pkgEnums,
		enums:   result,
	}
}

func newEnum(pkg *golang.Package, e Enum) Enum {
	return Enum{
		File:        pkg.AddGoFile(e.Name),
		Name:        e.Name,
		Description: e.Description,
		Type:        e.Type,
		Values:      e.Values,
	}
}

func (e Enums) GetEnumType(input string) (golang.IType, bool) {
	if strings.HasPrefix(input, specPrefix) {
		name := strings.TrimPrefix(input, specPrefix)
		result := golang.NewTypeMock(e.GetImport(), utils.Pascal(name), false, false)
		return result, true
	}
	return nil, false
}
