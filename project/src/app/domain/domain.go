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
	Service  service.Service
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

func newDomain(pkg *golang.Package, m *model.Model, enums *enum.Enums) Domain {
	pkgDomain := pkg.AddPackage(utils.Plural(m.Name))
	m.Build(pkgDomain, enums, "model")
	meta := model.Meta{
		PKG:   pkgDomain,
		Model: *m,
	}
	return Domain{
		pkg:      pkgDomain,
		Model:    *m,
		Repo:     repo.NewRepo("repo", meta),
		Service:  service.NewService("service", meta),
		Handlers: handlers.NewHandlers("handlers", meta),
	}
}
