package config

type ServiceOptions struct {
	FileNameTemplate           NameTemplate
	ExternalNameTemplate       NameTemplate
	InterfaceNameTemplate      NameTemplate
	ImplementationNameTemplate NameTemplate

	RepoFieldName string
}
