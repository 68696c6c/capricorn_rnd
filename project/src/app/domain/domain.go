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

type Map map[string]Domain

type Domain struct {
	pkg      *golang.Package
	Model    model.Model
	Repo     repo.Repo
	Service  *service.Service
	Handlers handlers.Handlers
}

func NewDomains(pkg *golang.Package, resources []*model.Model, enums *enum.Enums) Map {
	result := make(Map)
	for _, r := range resources {
		key := utils.Kebob(r.Name)
		result[key] = newDomain(pkg, r, enums)
	}
	return result
}

func newDomain(pkg *golang.Package, resource *model.Model, enums *enum.Enums) Domain {
	pkgDomain := pkg.AddPackage(utils.Plural(resource.Name))
	modelType := resource.Build(pkgDomain, enums, "model")
	actions := resource.Actions
	if len(actions) == 0 {
		actions = model.GetAllActions()
	}
	meta := model.Meta{
		ModelType: modelType,
		Actions:   actions,
	}
	domainRepo := repo.NewRepo(pkgDomain, "repo", meta)
	var domainService *service.Service
	if len(resource.Custom) > 0 {
		domainService = service.NewService(pkgDomain, "service", service.Meta{
			RepoType: domainRepo.GetInterfaceType(),
			Methods:  resource.Custom,
		})
	}
	return Domain{
		pkg:      pkgDomain,
		Model:    *resource,
		Repo:     domainRepo,
		Handlers: handlers.NewHandlers(pkgDomain, "handlers", meta),
		Service:  domainService,
	}
}
