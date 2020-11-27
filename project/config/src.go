package config

type SrcOptions struct {
	PkgName string

	App  AppOptions
	Cmd  CmdOptions
	Db   DbOptions
	Http HttpOptions
}

type AppOptions struct {
	PkgName string

	Enums                  EnumOptions
	Domain                 DomainOptions
	ServiceContainer       ServiceContainerOptions
	ServiceContainerConfig ServiceContainerConfigOptions
}

type ServiceContainerOptions struct {
	FileName      string
	TypeName      string
	SingletonName string

	DbArgName     string
	LoggerArgName string

	DbFieldName     string
	LoggerFieldName string
	ErrorsFieldName string

	AppInitFuncName string

	RepoVarNameTemplate NameTemplate
}

type ServiceContainerConfigOptions struct {
	FileName string
}

type DbOptions struct {
	PkgName    string
	Migrations MigrationsOptions
}

type MigrationsOptions struct {
	PkgName  string
	FileName string
}

type HttpOptions struct {
	PkgName string
	Routes  RoutesOptions
}

type RoutesOptions struct {
	FileName           string
	ApiGroupName       string
	GroupVarName       string
	ServicesArgName    string
	RouterInitFuncName string
}
