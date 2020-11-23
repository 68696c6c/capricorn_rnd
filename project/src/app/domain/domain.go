package domain

import (
	"strings"

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
	model               *model.Model
	repo                *repo.Repo
	service             *service.Service
	handlers            *handlers.RouteGroup
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
	domain := &Domain{
		Package: pkgApp.AddPackage(utils.Plural(resource.Name)),
	}

	domain.model = resource.Build(domain, enums, "model")

	var hasRepo bool
	var hasHandlers bool
	for _, a := range resource.Actions {
		if a == model.ActionNone {
			domain.hasRepo = false
			domain.hasHandlers = false
			return domain
		} else {
			hasRepo = true
		}
		if !strings.HasPrefix(string(a), "repo:") {
			hasHandlers = true
		}
	}
	if len(resource.Actions) == 0 {
		hasRepo = true
		hasHandlers = true
	}
	domain.hasRepo = hasRepo
	domain.hasHandlers = hasHandlers

	meta := model.Meta{
		ModelType:  *domain.model.GetType(),
		SingleName: utils.Singular(resource.Name),
		PluralName: utils.Plural(resource.Name),
		Actions:    domain.model.GetActions(),
	}

	repoFileName := "repo"
	domain.repo = repo.NewRepo(domain, repoFileName, meta)
	domain.externalRepoName = domain.Package.GetName() + "_" + repoFileName

	serviceFileName := "service"
	domain.service = service.NewService(domain, serviceFileName, service.Meta{
		RepoType: domain.repo.GetInterfaceType(),
		Methods:  resource.Custom,
	})
	domain.externalServiceName = domain.Package.GetName() + "_" + serviceFileName

	if domain.hasHandlers {
		domain.handlers = handlers.NewRouteGroup(domain, "handlers", meta, domain.repo)
	}

	return domain
}

func (d *Domain) HasRepo() bool {
	return d.hasRepo
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
	return d.hasHandlers
}

func (d *Domain) GetHandlers() *handlers.RouteGroup {
	return d.handlers
}

// TODO: avoid needing to do this somehow
func (d *Domain) SetHandlersErrorsRef(ref string) {
	d.handlers.SetErrorsRef(ref)
}

// TODO: avoid needing to do this somehow
func (d *Domain) SetHandlersRepoRef(ref string) {
	d.handlers.SetRepoRef(ref)
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
