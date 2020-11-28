package domain

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/handlers"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/repo"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/service"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
)

type Map map[string]*Domain

type Domains struct {
	domainMap   Map
	pkgModels   golang.IPackage
	pkgRepos    golang.IPackage
	pkgHandlers golang.IPackage
	pkgServices golang.IPackage
}

func (d *Domains) GetMap() Map {
	return d.domainMap
}

func (d *Domains) GetImportModels() string {
	return d.pkgModels.GetImport()
}

func (d *Domains) GetImportRepos() string {
	return d.pkgRepos.GetImport()
}

func (d *Domains) GetImportHandlers() string {
	return d.pkgHandlers.GetImport()
}

func (d *Domains) GetImportServices() string {
	return d.pkgServices.GetImport()
}

type Domain struct {
	model               *model.Model
	repo                *repo.Repo
	service             *service.Service
	handlers            *handlers.Handlers
	externalRepoName    string
	externalServiceName string
	hasRepo             bool
	hasHandlers         bool
}

func NewDomains(pkg golang.IPackage, o config.DomainOptions, resources []config.Model, enums *enum.Enums) *Domains {
	pkgModels := pkg.AddPackage(o.Model.PkgName)
	pkgRepos := pkg.AddPackage(o.Repo.PkgName)
	pkgHandlers := pkg.AddPackage(o.Handlers.PkgName)
	pkgServices := pkg.AddPackage(o.Service.PkgName)

	domainMap := make(Map)
	for _, r := range resources {
		d := newDomain(pkgModels, pkgRepos, pkgServices, pkgHandlers, o, r, enums)
		domainMap[r.Name] = d
	}
	return &Domains{
		domainMap:   domainMap,
		pkgModels:   pkgModels,
		pkgRepos:    pkgRepos,
		pkgHandlers: pkgHandlers,
		pkgServices: pkgServices,
	}
}

func newDomain(pkgModels, pkgRepos, pkgServices, pkgHandlers golang.IPackage, o config.DomainOptions, resource config.Model, enums *enum.Enums) *Domain {
	meta := config.NewDomainMeta(resource.Name, resource.Actions, resource.Custom, pkgModels, pkgRepos, pkgServices, pkgHandlers)

	domainModel := model.Build(pkgModels, o.Model, resource, enums)
	meta.SetModel(domainModel)

	domainRepo := repo.Build(pkgRepos, o.Repo, meta)

	var domainService *service.Service
	var domainHandlers *handlers.Handlers

	var externalRepoName string
	var externalServiceName string
	if domainRepo != nil {
		externalRepoName = domainRepo.GetExternalName()
		meta.SetRepo(domainRepo.GetInterfaceType())

		domainService = service.Build(pkgServices, o.Service, meta)
		if domainService != nil {
			externalServiceName = domainService.GetExternalName()
		}

		domainHandlers = handlers.Build(pkgHandlers, o.Handlers, meta)
	}

	return &Domain{
		model:               domainModel,
		repo:                domainRepo,
		externalRepoName:    externalRepoName,
		service:             domainService,
		externalServiceName: externalServiceName,
		handlers:            domainHandlers,
	}
}

func (d *Domain) GetModelType() golang.IType {
	return d.model
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

func (d *Domain) GetExternalRepoName() string {
	return d.externalRepoName
}

func (d *Domain) GetExternalServiceName() string {
	if !d.HasService() {
		return ""
	}
	return d.externalServiceName
}
