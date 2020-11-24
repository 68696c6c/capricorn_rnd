package domain

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
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
	model               *model.Model
	repo                *repo.Repo
	service             *service.Service
	handlers            *handlers.Group
	externalRepoName    string
	externalServiceName string
	hasRepo             bool
	hasHandlers         bool
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
	domainPkgName := utils.Plural(resource.Name)
	domain := &Domain{
		Package: pkgApp.AddPackage(domainPkgName),
	}

	domain.model = resource.Build(domain, enums, "model")

	meta := config.NewDomainResource(resource.Name, resource.Actions, resource.Custom)
	meta.SetModel(domain.model.GetType())

	repoFileName := "repo"
	domainRepo := repo.NewRepo(domain, repoFileName, meta)
	if domainRepo != nil {
		domain.repo = domainRepo
		domain.externalRepoName = domainPkgName + "_" + repoFileName
		meta.SetRepo(domain.repo.GetInterfaceType())

		serviceFileName := "service"
		domain.service = service.NewService(domain, serviceFileName, meta)
		domain.externalServiceName = domainPkgName + "_" + serviceFileName

		domain.handlers = handlers.NewGroup(domain, "handlers", meta)
	}

	return domain
}

func (d *Domain) HasRepo() bool {
	return d.repo != nil
}

func (d *Domain) GetRepoConstructor() *golang.Function {
	return d.repo.GetConstructor()
}

func (d *Domain) GetRepoInterfaceType() golang.IType {
	return d.repo.GetInterfaceType()
}

func (d *Domain) HasService() bool {
	return d.service != nil
}

func (d *Domain) GetServiceConstructor() *golang.Function {
	return d.service.GetConstructor()
}

func (d *Domain) GetServiceInterfaceType() golang.IType {
	return d.service.GetInterfaceType()
}

func (d *Domain) HasHandlers() bool {
	return d.handlers != nil
}

func (d *Domain) GetHandlers() *handlers.Group {
	return d.handlers
}

// Returns the name of the field this repo should live under in the service container.
// For DDD apps, this will be the package name + the repo name, for MVC apps it is just the repo name.
func (d *Domain) GetExternalRepoName() string {
	return d.externalRepoName
}

func (d *Domain) GetExternalServiceName() string {
	if !d.HasService() {
		return ""
	}
	return d.externalServiceName
}
