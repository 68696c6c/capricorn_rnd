package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Project struct {
	Name      string     `yaml:"name,omitempty"`
	Module    string     `yaml:"module,omitempty"`
	License   string     `yaml:"license,omitempty"`
	Author    AuthorMeta `yaml:"author,omitempty"`
	Ops       Ops        `yaml:"ops"`
	Commands  []Command  `yaml:"commands"`
	Enums     []Enum     `yaml:"enums"`
	Resources []Model    `yaml:"resources"`
}

type AuthorMeta struct {
	Name         string `yaml:"name,omitempty"`
	Email        string `yaml:"email,omitempty"`
	Organization string `yaml:"organization,omitempty"`
}

func NewProjectFromSpec(specPath string) (*Project, error) {
	file, err := ioutil.ReadFile(specPath)
	if err != nil {
		return &Project{}, errors.Wrap(err, "failed to read spec file")
	}

	result := Project{}
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return &Project{}, errors.Wrap(err, "failed to unmarshal spec")
	}

	return &result, nil
}

// Project options should be all the internal, non-user-provided configuration.
// The idea is that you could swap different ProjectOptions to generate different implementations of the same user-
// provided project spec.
type ProjectOptions struct {
	BasePath string
	Src      SrcOptions
}

func NewProjectOptions(basePath string) ProjectOptions {
	repoPaginationFuncName := "ApplyPaginationToQuery"
	repoFilterFuncName := "Filter"
	repoDeleteFuncName := "Delete"
	repoGetByIdFuncName := "GetById"
	repoSaveFuncName := "Save"
	return ProjectOptions{
		BasePath: basePath,

		Src: SrcOptions{
			PkgName: "src",

			App: AppOptions{
				PkgName: "app",

				Enums: EnumOptions{
					PkgName:    "enums",
					FileName:   "enums",
					SpecPrefix: "enum:",

					InputArgName: "input",

					FromStringFuncNameSuffix: "FromString",
					ScanFuncName:             "Scan",
					ValueFuncName:            "Value",
				},

				Domain: DomainOptions{
					Model: ModelOptions{
						PkgName:          "models",
						FileNameTemplate: nameTemplateResourceSingular,
						TypeNameTemplate: nameTemplateResourceSingular,
					},
					Repo: RepoOptions{
						PkgName:                    "repos",
						FileNameTemplate:           NameTemplateF("%s_repo", nameTemplateResourcePlural),
						ExternalNameTemplate:       NameTemplateF("%s_repo", nameTemplateResourcePlural),
						InterfaceNameTemplate:      NameTemplateF("%s_repo", nameTemplateResourcePlural),
						ImplementationNameTemplate: NameTemplateF("%s_repo_gorm", nameTemplateResourcePlural),

						ModelArgName: "model",
						QueryArgName: "query",
						DbArgName:    "dbConnection",
						IdArgName:    "id",

						DbFieldName: "db",

						BaseQueryFuncName:   "getBaseQuery",
						PaginationFuncName:  repoPaginationFuncName,
						FilterFuncName:      repoFilterFuncName,
						FilterQueryFuncName: "getFilteredQuery",
						DeleteFuncName:      repoDeleteFuncName,
						GetByIdFuncName:     repoGetByIdFuncName,
						SaveFuncName:        repoSaveFuncName,
					},
					Service: ServiceOptions{
						PkgName:                    "services",
						FileNameTemplate:           NameTemplateF("%s_service", nameTemplateResourcePlural),
						ExternalNameTemplate:       NameTemplateF("%s_service", nameTemplateResourcePlural),
						InterfaceNameTemplate:      NameTemplateF("%s_service", nameTemplateResourcePlural),
						ImplementationNameTemplate: NameTemplateF("%s_service_implementation", nameTemplateResourcePlural),
						RepoFieldName:              "repo",
					},
					Handlers: HandlersOptions{
						PkgName:          "handlers",
						FileNameTemplate: NameTemplateF("%s", nameTemplateResourcePlural),
						UriTemplate:      NameTemplateF("%s", nameTemplateResourcePlural),

						RepoPaginationFuncName: repoPaginationFuncName,
						RepoFilterFuncName:     repoFilterFuncName,
						RepoGetByIdFuncName:    repoGetByIdFuncName,
						RepoDeleteFuncName:     repoDeleteFuncName,
						RepoSaveFuncName:       repoSaveFuncName,

						ParamNameId:         "id",
						ContextArgName:      "c",
						ErrorsArgName:       "errorHandler",
						RepoArgNameTemplate: NameTemplateF("%s_repo", nameTemplateResourcePlural),

						CreateRequestNameTemplate:    NameTemplateF("create_%s_request", nameTemplateResourceSingular),
						UpdateRequestNameTemplate:    NameTemplateF("update_%s_request", nameTemplateResourceSingular),
						ResourceResponseNameTemplate: NameTemplateF("%s_response", nameTemplateResourceSingular),
						ListResponseNameTemplate:     NameTemplateF("%s_response", nameTemplateResourcePlural),

						CreateNameTemplate: NameTemplateF("create_%s", nameTemplateResourceSingular),
						UpdateNameTemplate: NameTemplateF("update_%s", nameTemplateResourceSingular),
						ViewNameTemplate:   NameTemplateF("view_%s", nameTemplateResourceSingular),
						ListNameTemplate:   NameTemplateF("list_%s", nameTemplateResourcePlural),
						DeleteNameTemplate: NameTemplateF("delete_%s", nameTemplateResourceSingular),
					},
				},

				ServiceContainer: ServiceContainerOptions{
					FileName:            "app",
					TypeName:            "ServiceContainer",
					SingletonName:       "container",
					DbArgName:           "dbConnection",
					LoggerArgName:       "logger",
					DbFieldName:         "DB",
					LoggerFieldName:     "Logger",
					ErrorsFieldName:     "Errors",
					AppInitFuncName:     "GetApp",
					RepoVarNameTemplate: NameTemplateF("%s_repo", nameTemplateResourcePlural),
				},
				ServiceContainerConfig: ServiceContainerConfigOptions{
					FileName: "config",
				},
			},

			Cmd: CmdOptions{
				PkgName: "cmd",

				CmdArgName:  "cmd",
				ArgsArgName: "args",

				RootCommandUse: "app",

				ServerFileName:   "server",
				ServerCommandUse: "server",

				MigrateFileName:   "migrate",
				MigrateCommandUse: "migrate",
			},

			Db: DbOptions{
				PkgName: "db",
				Migrations: MigrationsOptions{
					PkgName: "migrations",
					// @TODO: use timestamps and think about how to handle initial migrations with AutoMigrate safely...
					InitialMigrationTimestamp: "20201127164400",
					InitialMigrationName:      "initial_migration",
				},
			},

			Http: HttpOptions{
				PkgName: "http",
				Routes: RoutesOptions{
					FileName:           "routes",
					ApiGroupName:       "api",
					GroupVarName:       "g",
					ServicesArgName:    "s",
					RouterInitFuncName: "InitRouter",
				},
			},
		},
	}
}
