package config

type RepoOptions struct {
	PkgName                    string
	FileNameTemplate           NameTemplate
	ExternalNameTemplate       NameTemplate
	InterfaceNameTemplate      NameTemplate
	ImplementationNameTemplate NameTemplate

	ModelArgName string
	QueryArgName string
	DbArgName    string
	IdArgName    string

	DbFieldName string

	BaseQueryFuncName   string
	PaginationFuncName  string
	FilterFuncName      string
	FilterQueryFuncName string
	DeleteFuncName      string
	GetByIdFuncName     string
	SaveFuncName        string
}
