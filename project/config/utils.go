package config

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Action string

const (
	ActionNone   Action = "none"
	ActionCreate Action = "create"
	ActionView   Action = "view"
	ActionList   Action = "list"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"

	ActionRepoCreate Action = "repo:create"
	ActionRepoView   Action = "repo:view"
	ActionRepoList   Action = "repo:list"
	ActionRepoUpdate Action = "repo:update"
	ActionRepoDelete Action = "repo:delete"
)

type ProjectOptions struct {
	RepoPaginationFuncName string
	RepoFilterFuncName     string

	HandlersFileName string
	RepoFileName     string
	ServiceFileName  string
	ModelFileName    string
}

type DomainResource struct {
	RepoType      golang.IType
	RepoFieldName string

	// ServiceType      golang.IType
	// ServiceFieldName string

	ModelType golang.IType
	ModelName string

	NameSingular string
	NamePlural   string
	// NameFirstLetter string
	RepoActions    []Action
	HandlerActions []Action
	ServiceActions []string

	RepoPaginationFuncName string
	RepoFilterFuncName     string
}

func removeDuplicateActions(actions []Action) []Action {
	keys := make(map[Action]bool)
	var result []Action
	for _, i := range actions {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			result = append(result, i)
		}
	}
	return result
}

func getDefaultActions() []Action {
	return []Action{ActionCreate, ActionView, ActionList, ActionUpdate, ActionDelete}
}

func getResourceActions(resourceActions []Action) (repoActions, handlerActions []Action) {
	if len(resourceActions) == 0 {
		resourceActions = getDefaultActions()
	}
	for _, a := range resourceActions {
		if a == ActionNone {
			return []Action{}, []Action{}
		}
		switch a {
		case ActionRepoCreate:
			repoActions = append(repoActions, ActionCreate)
			break
		case ActionRepoView:
			repoActions = append(repoActions, ActionView)
			break
		case ActionRepoList:
			repoActions = append(repoActions, ActionList)
			break
		case ActionRepoUpdate:
			repoActions = append(repoActions, ActionUpdate)
			break
		case ActionRepoDelete:
			repoActions = append(repoActions, ActionDelete)
			break
		default:
			handlerActions = append(handlerActions, a)
			repoActions = append(repoActions, a)
		}
	}
	return removeDuplicateActions(repoActions), handlerActions
}

func NewDomainResource(resourceName string, resourceActions []Action, customActions []string) *DomainResource {
	repoActions, handlerActions := getResourceActions(resourceActions)
	return &DomainResource{
		RepoType:               nil,
		ModelType:              nil,
		NameSingular:           utils.Singular(resourceName),
		NamePlural:             utils.Plural(resourceName),
		RepoActions:            repoActions,
		HandlerActions:         handlerActions,
		ServiceActions:         customActions,
		RepoPaginationFuncName: "ApplyPaginationToQuery",
		RepoFilterFuncName:     "Filter",
	}
}

func (r *DomainResource) SetModel(modelType golang.IType) {
	r.ModelType = modelType
}

func (r *DomainResource) SetRepo(repoType golang.IType) {
	r.RepoType = repoType
}

func (r *DomainResource) GetModelType() *golang.Type {
	return r.ModelType.CopyType()
}

func (r *DomainResource) GetRepoType() *golang.Type {
	return r.RepoType.CopyType()
}

func (r *DomainResource) GetModelName() string {
	return r.ModelName
}

func (r *DomainResource) GetRepoActions() []Action {
	return r.RepoActions
}

func (r *DomainResource) GetHandlerActions() []Action {
	return r.HandlerActions
}

func (r *DomainResource) GetServiceActions() []string {
	return r.ServiceActions
}
