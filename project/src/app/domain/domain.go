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
	repo                *repo.Repo
	service             *service.Service
	handlers            *handlers.Handlers
	externalRepoName    string
	externalServiceName string
	hasRepo             bool
	hasHandlers         bool
}

func NewDomains(pkg golang.IPackage, o config.DomainOptions, resources []config.Model, enums *enum.Enums) Map {
	result := make(Map)
	for _, r := range resources {
		d := newDomain(pkg, o, r, enums)
		result[r.Name] = d
	}
	return result
}

func newDomain(pkg golang.IPackage, o config.DomainOptions, resource config.Model, enums *enum.Enums) *Domain {
	domainPkgName := utils.Plural(resource.Name)
	pkgDomain := pkg.AddPackage(domainPkgName)
	meta := config.NewDomainMeta(resource.Name, resource.Actions, resource.Custom)

	domainModel := model.Build(pkgDomain, o.Model, resource, enums)
	meta.SetModel(domainModel)

	domainRepo := repo.Build(pkgDomain, o.Repo, meta)

	var domainService *service.Service
	var domainHandlers *handlers.Handlers

	var externalRepoName string
	var externalServiceName string
	if domainRepo != nil {
		externalRepoName = domainRepo.GetExternalName()
		meta.SetRepo(domainRepo.GetInterfaceType())

		domainService = service.Build(pkgDomain, o.Service, meta)
		if domainService != nil {
			externalServiceName = domainService.GetExternalName()
		}

		domainHandlers = handlers.Build(pkgDomain, o.Handlers, meta)
	}

	return &Domain{
		Package:             pkgDomain,
		repo:                domainRepo,
		externalRepoName:    externalRepoName,
		service:             domainService,
		externalServiceName: externalServiceName,
		handlers:            domainHandlers,
	}
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

func (d *Domain) GetHandlers() *handlers.Handlers {
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
