package model

import (
	"path"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const dddModelName = "Model"

type Meta struct {
	PKG *golang.Package
	Model
}

func getAssumedModelType(baseImport, inputName string, isPointer bool) golang.IType {
	pkgName := utils.Plural(inputName)
	return golang.NewStructFromType(golang.Type{
		Import:    path.Join(baseImport, pkgName),
		Package:   pkgName,
		Name:      dddModelName,
		IsPointer: isPointer,
		IsSlice:   false,
	})
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

func newBaseModelStruct() *golang.Struct {
	return golang.NewStructFromType(golang.Type{
		Import:    golang.ImportGoat,
		Package:   "goat",
		Name:      dddModelName,
		IsPointer: false,
		IsSlice:   false,
	})
}

func getModelSoftDelete() *golang.Struct {
	result := newBaseModelStruct()
	result.AddField(makeField("id", getIdType(), true))
	result.AddField(makeField("created_at", getTimeType(false), true))
	result.AddField(makeField("updated_at", getTimeType(true), true))
	result.AddField(makeField("deleted_at", getTimeType(true), true))
	return result
}

func getModelHardDelete() *golang.Struct {
	result := newBaseModelStruct()
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
