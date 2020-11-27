package config

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
