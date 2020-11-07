package enum

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Map map[string]Enum

type Enums struct {
	pkg   *golang.Package
	enums Map
}

type Enum struct {
	file        *golang.File
	Name        string   `yaml:"name,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Values      []string `yaml:"values,omitempty"`
}

func NewEnums(pkg *golang.Package, enums []Enum) Enums {
	result := make(Map)
	pkgEnums := pkg.AddPackage("enums")
	for _, e := range enums {
		key := utils.Kebob(e.Name)
		result[key] = newEnum(pkgEnums, e)
	}
	return Enums{
		pkg:   pkgEnums,
		enums: result,
	}
}

func newEnum(pkg *golang.Package, e Enum) Enum {
	return Enum{
		file:        pkg.AddGoFile(e.Name),
		Name:        e.Name,
		Description: e.Description,
		Type:        e.Type,
		Values:      e.Values,
	}
}
