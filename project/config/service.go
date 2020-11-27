package config

type ServiceOptions struct {
	PkgName                    string
	FileNameTemplate           NameTemplate
	ExternalNameTemplate       NameTemplate
	InterfaceNameTemplate      NameTemplate
	ImplementationNameTemplate NameTemplate

	RepoFieldName string
}
