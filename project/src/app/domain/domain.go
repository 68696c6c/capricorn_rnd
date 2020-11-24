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
	handlers            *handlers.Handlers
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
	pkgDomain := pkgApp.AddPackage(domainPkgName)
	meta := config.NewDomainMeta(resource.Name, resource.Actions, resource.Custom)

	domainModel := resource.Build(pkgDomain, enums, "model")
	meta.SetModel(domainModel.GetType())

	repoFileName := "repo"
	domainRepo := repo.Build(pkgDomain, repoFileName, meta)

	serviceFileName := "service"
	var domainService *service.Service
	var domainHandlers *handlers.Handlers

	if domainRepo != nil {
		meta.SetRepo(domainRepo.GetInterfaceType())

		domainService = service.Build(pkgDomain, serviceFileName, meta)

		domainHandlers = handlers.Build(pkgDomain, "handlers", meta)
	}

	return &Domain{
		Package:             pkgDomain,
		model:               domainModel,
		repo:                domainRepo,
		externalRepoName:    domainPkgName + "_" + repoFileName,
		service:             domainService,
		externalServiceName: domainPkgName + "_" + serviceFileName,
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
