package domain

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/handlers"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/repo"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/service"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Map map[string]*Domain

type Domain struct {
	*golang.Package
	Model               *model.Model
	Repo                *repo.Repo
	Service             *service.Service
	Handlers            handlers.Handlers
	externalRepoName    string
	externalServiceName string
}

func NewDomains(pkgApp *golang.Package, resources []*model.Model, enums *enum.Enums) Map {
	result := make(Map)
	for _, r := range resources {
		d := newDomain(pkgApp, r, enums)
		result[d.Package.GetName()] = d
	}
	return result
}

func newDomain(pkgApp *golang.Package, resource *model.Model, enums *enum.Enums) *Domain {
	domain := &Domain{
		Package: pkgApp.AddPackage(utils.Plural(resource.Name)),
	}

	domain.Model = resource.Build(domain, enums, "model")

	meta := model.Meta{
		ModelType: *domain.Model.GetType(),
		Actions:   domain.Model.GetActions(),
	}

	repoFileName := "repo"
	domain.Repo = repo.NewRepo(domain, repoFileName, meta)
	domain.externalRepoName = domain.Package.GetName() + "_" + repoFileName

	serviceFileName := "service"
	domain.Service = service.NewService(domain, serviceFileName, service.Meta{
		RepoType: domain.Repo.GetInterfaceType(),
		Methods:  resource.Custom,
	})
	domain.externalServiceName = domain.Package.GetName() + "_" + serviceFileName

	domain.Handlers = handlers.NewHandlers(domain, "handlers", meta)

	return domain
}

// Returns the name of the field this repo should live under in the service container.
// For DDD apps, this will be the package name + the repo name, for MVC apps it is just the repo name.
func (d *Domain) GetExternalRepoName() string {
	return d.externalRepoName
}

func (d *Domain) GetExternalServiceName() string {
	if d.Service == nil {
		return ""
	}
	return d.externalServiceName
}
