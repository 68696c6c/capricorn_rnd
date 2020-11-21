package container

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
	"github.com/pkg/errors"
)

type containerMeta struct {
	domains         domain.Map
	structType      *golang.Struct
	singletonName   string
	dbFieldName     string
	loggerFieldName string
	errorsFieldName string
}

type domainServices struct {
	repo    *golang.Field
	service *golang.Field
}

type Container struct {
	*golang.File
	containerType *golang.Struct
	errorsField   *golang.Field
	fields        map[string]domainServices
}

func NewContainer(pkg *golang.Package, domains domain.Map) *Container {
	containerStruct := golang.NewStruct("ServiceContainer", false, false)
	meta := containerMeta{
		structType:      containerStruct,
		singletonName:   "container",
		domains:         domains,
		dbFieldName:     "DB",
		loggerFieldName: "Logger",
		errorsFieldName: "Errors",
	}

	containerStruct.AddConstructor(makeConstructor(meta))

	containerStruct.AddField(golang.NewField(meta.dbFieldName, goat.MakeTypeDbConnection(), true))
	containerStruct.AddField(golang.NewField(meta.loggerFieldName, goat.MakeTypeLogger(), true))

	errorsField := golang.NewField(meta.errorsFieldName, goat.MakeTypeErrorHandler(), true)
	containerStruct.AddField(errorsField)

	result := &Container{
		File:          pkg.AddGoFile("app"),
		containerType: containerStruct,
		errorsField:   errorsField,
		fields:        make(map[string]domainServices),
	}

	for domainKey, d := range domains {
		repoType := d.Repo.GetInterfaceType()
		repoField := golang.NewField(d.GetExternalRepoName(), repoType, true)
		containerStruct.AddField(repoField)

		services := domainServices{
			repo: repoField,
		}

		if d.Service != nil {
			serviceType := d.Service.GetInterfaceType()
			serviceField := golang.NewField(d.GetExternalServiceName(), serviceType, true)
			containerStruct.AddField(serviceField)
			services.service = serviceField
		}

		result.fields[domainKey] = services
	}
	result.AddVar(golang.NewVar(meta.singletonName, "", containerStruct.CopyType(), false))
	result.AddStruct(containerStruct)

	return result
}

func (c *Container) GetContainerType() *golang.Struct {
	return c.containerType
}

func (c *Container) ErrorHandlerField() *golang.Field {
	return c.errorsField
}

func (c *Container) GetDomainRepoField(domainKey string) (*golang.Field, error) {
	result, ok := c.fields[domainKey]
	if !ok {
		return nil, errors.Errorf("unexpected domain key '%'", domainKey)
	}
	if result.repo == nil {
		return nil, errors.New("domain does not have a repo")
	}
	return result.repo, nil
}
