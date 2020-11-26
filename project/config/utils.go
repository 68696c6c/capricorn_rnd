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

	// @TODO use project options instead
	AppInitFuncName    = "InitApp"
	RouterInitFuncName = "InitRouter"
)

type AuthorMeta struct {
	Name         string `yaml:"name,omitempty"`
	Email        string `yaml:"email,omitempty"`
	Organization string `yaml:"organization,omitempty"`
}

// @TODO: just keeping track of things here, but need to refactor to actually use this struct and pass all configuration down from the top instead of strewing it around everywhere.
// Project options should be all the internal, non-user-provided configuration.
// The idea is that you could swap different ProjectOptions to generate different implementations of the same user-
// provided project spec, i.e. generate a DDD or an MVC variants of the same project.
type ProjectOptions struct {
	// RepoPaginationFuncName string
	// RepoFilterFuncName     string
	//
	// HandlersFileName   string
	// RepoFileName       string
	// ServiceFileName    string
	// ModelFileName      string
	// CmdRootFileName    string
	// CmdServerFileName  string
	// CmdMigrateFileName string
	//
	// CmdRootVarName string
	//
	// AppInitFuncName    string
	// RouterInitFuncName string
}

type CmdMeta struct {
	RootVarName    string
	RootCommandUse string
	ProjectName    string
	License        string
	Author         AuthorMeta

	RootFileName    string
	ServerFileName  string
	MigrateFileName string
}

func NewCmdMeta(projectName, license string, author AuthorMeta) *CmdMeta {
	return &CmdMeta{
		RootVarName:     "Root",
		RootCommandUse:  "app",
		ProjectName:     projectName,
		License:         license,
		Author:          author,
		RootFileName:    "root",
		ServerFileName:  "server",
		MigrateFileName: "migrate",
	}
}

type DomainMeta struct {
	NameSingular string
	NamePlural   string

	ModelType golang.IType
	ModelName string

	RepoType               golang.IType
	RepoFieldName          string
	RepoPaginationFuncName string
	RepoFilterFuncName     string

	RepoActions    []Action
	HandlerActions []Action
	ServiceActions []string
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

func NewDomainMeta(resourceName string, resourceActions []Action, customActions []string) *DomainMeta {
	repoActions, handlerActions := getResourceActions(resourceActions)
	return &DomainMeta{
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

func (r *DomainMeta) SetModel(modelType golang.IType) {
	r.ModelType = modelType
}

func (r *DomainMeta) SetRepo(repoType golang.IType) {
	r.RepoType = repoType
}

func (r *DomainMeta) GetModelType() *golang.Type {
	return r.ModelType.CopyType()
}

func (r *DomainMeta) GetRepoType() *golang.Type {
	return r.RepoType.CopyType()
}

func (r *DomainMeta) GetModelName() string {
	return r.ModelName
}

func (r *DomainMeta) GetRepoActions() []Action {
	return r.RepoActions
}

func (r *DomainMeta) GetHandlerActions() []Action {
	return r.HandlerActions
}

func (r *DomainMeta) GetServiceActions() []string {
	return r.ServiceActions
}
