package model

import (
	"path"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const dddModelName = "Model"

func getAssumedDDDModelType(baseImport, inputName string, isPointer bool) golang.IType {
	pkgName := utils.Plural(inputName)
	imp := path.Join(baseImport, pkgName)
	result := golang.MockStruct(imp, dddModelName, isPointer, false)
	return result
}
