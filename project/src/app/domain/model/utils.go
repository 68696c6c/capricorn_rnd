package model

import (
	"path"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const dddModelName = "Model"

type Action string

const (
	ActionNone       = "none"
	ActionCreate     = "create"
	ActionView       = "view"
	ActionList       = "list"
	ActionUpdate     = "update"
	ActionDelete     = "delete"
	ActionRepoCreate = "repo:create"
	ActionRepoView   = "repo:view"
	ActionRepoList   = "repo:list"
	ActionRepoUpdate = "repo:update"
	ActionRepoDelete = "repo:delete"
)

func GetAllActions() []Action {
	return []Action{ActionCreate, ActionView, ActionList, ActionUpdate, ActionDelete}
}

type Meta struct {
	ModelType  Type
	SingleName string
	PluralName string
	Actions    []Action
}

func getAssumedDDDModelType(baseImport, inputName string, isPointer bool) golang.IType {
	pkgName := utils.Plural(inputName)
	imp := path.Join(baseImport, pkgName)
	result := golang.MockStruct(imp, dddModelName, isPointer, false)
	return result
}
