package config

type HandlersOptions struct {
	FileNameTemplate NameTemplate
	UriTemplate      NameTemplate

	RepoPaginationFuncName string
	RepoFilterFuncName     string
	RepoGetByIdFuncName    string
	RepoSaveFuncName       string
	RepoDeleteFuncName     string

	ParamNameId         string
	ContextArgName      string
	ErrorsArgName       string
	RepoArgNameTemplate NameTemplate

	CreateRequestNameTemplate    NameTemplate
	UpdateRequestNameTemplate    NameTemplate
	ResourceResponseNameTemplate NameTemplate
	ListResponseNameTemplate     NameTemplate

	CreateNameTemplate NameTemplate
	UpdateNameTemplate NameTemplate
	ViewNameTemplate   NameTemplate
	ListNameTemplate   NameTemplate
	DeleteNameTemplate NameTemplate
}
