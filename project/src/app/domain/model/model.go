package model

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const dddModelName = "Model"

type Meta struct {
	PKG   *golang.Package
	Model Model
}

type fieldMap map[string]golang.Field

type fieldCategories struct {
	all      fieldMap `desc:"all fields, including base model fields, relations, and user defined fields"`
	model    fieldMap `desc:"fields that are defined on the model struct. includes relations but not base model fields"`
	database fieldMap `desc:"fields that exist in the database.  includes base model fields but not relations"`
}

func (f fieldCategories) AddAllField(key string, field golang.Field) {
	k := utils.Snake(key)
	f.all[k] = field
	f.model[k] = field
	f.database[k] = field
}

func (f fieldCategories) AddModelField(key string, field golang.Field) {
	k := utils.Snake(key)
	f.all[k] = field
	f.model[k] = field
}

func (f fieldCategories) AddDbField(key string, field golang.Field) {
	k := utils.Snake(key)
	f.all[k] = field
	f.database[k] = field
}

type Model struct {
	*golang.File `yaml:"-"`
	Name         string   `yaml:"name,omitempty"`
	Delete       string   `yaml:"delete,omitempty"`
	BelongsTo    []string `yaml:"belongs_to,omitempty"`
	HasMany      []string `yaml:"has_many,omitempty"`
	Fields       []Field  `yaml:"fields,omitempty"`
	Actions      []string `yaml:"actions,omitempty"`
	Custom       []string `yaml:"custom,omitempty"`
}

type Field struct {
	Name     string `yaml:"name,omitempty"`
	Type     string `yaml:"type,omitempty"`
	Enum     string `yaml:"enum,omitempty"`
	Required bool   `yaml:"required,omitempty"`
	Unique   bool   `yaml:"unique,omitempty"`
	Indexed  bool   `yaml:"indexed,omitempty"`
}

func (f Field) GetType() golang.IType {
	eType, isEnum := enum.GetEnumType(f.Type)
	if isEnum {
		return eType
	}
	return golang.NewTypeFromReference(f.Type)
}

func NewModel(fileName string, meta Meta) Model {
	result := &Model{
		File:      meta.PKG.AddGoFile(fileName),
		Name:      meta.Model.Name,
		BelongsTo: meta.Model.BelongsTo,
		HasMany:   meta.Model.HasMany,
		Fields:    meta.Model.Fields,
		Actions:   meta.Model.Actions,
		Custom:    meta.Model.Custom,
	}
	result.build()
	return *result
}

func (m *Model) build() {
	modelType := getModelSoftDelete()
	if m.Delete == "hard" {
		modelType = getModelHardDelete()
	}

	fields := fieldCategories{
		all:      make(fieldMap),
		model:    make(fieldMap),
		database: make(fieldMap),
	}

	// Base model fields.
	fields.AddModelField("model", golang.Field{
		Type: modelType,
	})
	for _, f := range modelType.GetStructFields() {
		fields.AddDbField(f.Name, f)
	}

	// User defined fields.
	for _, f := range m.Fields {
		fields.AddAllField(f.Name, makeField(f.Name, f.GetType(), true))
	}

	// Has many.
	for _, relation := range m.HasMany {
		name := utils.Pascal(relation)
		relModel := getAssumedModelType(relation, true)
		t := golang.MakeSliceType(false, relModel)
		fields.AddModelField(name, makeField(name, t, true))
	}

	// Belongs to.
	for _, relation := range m.BelongsTo {
		// This is the field GORM will hydrate the relation in to.
		relationName := utils.Pascal(relation)
		relModel := getAssumedModelType(relation, true)
		fields.AddModelField(relationName, makeField(relationName, relModel, true))

		// This is the foreign key field that establishes the relation.
		relationIdName := utils.Pascal(relation + "_id")
		fields.AddAllField(relationIdName, makeField(relationIdName, getIdType(), true))
	}

	// Build the struct.
	modelStruct := &golang.Struct{
		Type: golang.Type{
			Import:    m.PKG.ImportPath,
			Package:   m.PKG.Name,
			Name:      dddModelName,
			IsPointer: false,
			IsSlice:   false,
		},
	}
	for _, f := range fields.model {
		println("adding field " + f.Name)
		modelStruct.AddField(f)
	}

	m.AddStruct(modelStruct)
}

func getAssumedModelType(pluralResourceName string, isPointer bool) golang.IType {
	return &golang.Struct{
		Type: golang.Type{
			Import:    "???",
			Package:   pluralResourceName,
			Name:      dddModelName,
			IsPointer: isPointer,
			IsSlice:   false,
		},
	}
}

func makeField(name string, t golang.IType, isExported bool) golang.Field {
	fieldName := utils.Camel(name)
	if isExported {
		fieldName = utils.Pascal(name)
	}
	return golang.Field{
		Name: fieldName,
		Type: t,
		Tags: []golang.Tag{
			{
				Key:    "json",
				Values: []string{utils.Snake(name)},
			},
		},
	}
}

func getModelSoftDelete() *golang.Struct {
	result := &golang.Struct{
		Type: golang.Type{
			Import:    golang.ImportGoat,
			Package:   "goat",
			Name:      dddModelName,
			IsPointer: false,
			IsSlice:   false,
		},
	}
	result.AddField(makeField("id", getIdType(), true))
	result.AddField(makeField("created_at", getTimeType(false), true))
	result.AddField(makeField("updated_at", getTimeType(true), true))
	result.AddField(makeField("deleted_at", getTimeType(true), true))
	return result
}

func getModelHardDelete() *golang.Struct {
	result := &golang.Struct{
		Type: golang.Type{
			Import:    golang.ImportGoat,
			Package:   "goat",
			Name:      dddModelName,
			IsPointer: false,
			IsSlice:   false,
		},
	}
	result.AddField(makeField("id", getIdType(), true))
	result.AddField(makeField("created_at", getTimeType(false), true))
	result.AddField(makeField("updated_at", getTimeType(true), true))
	return result
}

func getIdType() golang.IType {
	return golang.Type{
		Import:    golang.ImportGoat,
		Package:   "goat",
		Name:      "ID",
		IsPointer: false,
		IsSlice:   false,
	}
}

func getStringType() golang.IType {
	return golang.Type{
		Import:    "",
		Package:   "",
		Name:      "string",
		IsPointer: false,
		IsSlice:   false,
	}
}

func getTimeType(isPointer bool) golang.IType {
	return golang.Type{
		Import:    "time",
		Package:   "time",
		Name:      "Time",
		IsPointer: isPointer,
		IsSlice:   false,
	}
}
