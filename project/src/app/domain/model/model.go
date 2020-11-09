package model

import "github.com/68696c6c/capricorn_rnd/golang"

type Meta struct {
	PKG   *golang.Package
	Model Model
}

type Model struct {
	*golang.File `yaml:"-"`
	Name         string                  `yaml:"name,omitempty"`
	Delete       string                  `yaml:"delete,omitempty"`
	BelongsTo    []string                `yaml:"belongs_to,omitempty"`
	HasMany      []string                `yaml:"has_many,omitempty"`
	Fields       []*Field                `yaml:"fields,omitempty"`
	Actions      []string                `yaml:"actions,omitempty"`
	Custom       []string                `yaml:"custom,omitempty"`
	fields       map[string]golang.Field `desc:"built fields grouped by category"`
}

type Field struct {
	Name      string `yaml:"name,omitempty"`
	Type      string `yaml:"type,omitempty"`
	Enum      string `yaml:"enum,omitempty"`
	Required  bool   `yaml:"required,omitempty"`
	Unique    bool   `yaml:"unique,omitempty"`
	Indexed   bool   `yaml:"indexed,omitempty"`
	isPrimary bool
}

func NewModel(fileName string, meta Meta) Model {
	return Model{
		File:      meta.PKG.AddGoFile(fileName),
		Name:      meta.Model.Name,
		BelongsTo: meta.Model.BelongsTo,
		HasMany:   meta.Model.HasMany,
		Fields:    meta.Model.Fields,
		Actions:   meta.Model.Actions,
		Custom:    meta.Model.Custom,
	}
}
