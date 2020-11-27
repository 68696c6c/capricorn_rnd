package config

type HandlersOptions struct {
	PkgName          string
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
