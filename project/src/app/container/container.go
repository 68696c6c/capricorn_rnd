package container

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"

	"github.com/pkg/errors"
)

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

func NewContainer(pkg golang.IPackage, o config.ServiceContainerOptions, domains domain.Map) *Container {
	containerStruct := golang.NewStruct(o.TypeName, false, false)

	containerStruct.AddConstructor(makeConstructor(o, containerStruct, domains))

	containerStruct.AddField(golang.NewField(o.DbFieldName, goat.MakeTypeDbConnection(), true))
	containerStruct.AddField(golang.NewField(o.LoggerFieldName, goat.MakeTypeLogger(), true))

	errorsField := golang.NewField(o.ErrorsFieldName, goat.MakeTypeErrorHandler(), true)
	containerStruct.AddField(errorsField)

	result := &Container{
		File:          pkg.AddGoFile(o.FileName),
		containerType: containerStruct,
		errorsField:   errorsField,
		fields:        make(map[string]domainServices),
	}

	for resourceName, d := range domains {
		if !d.HasRepo() {
			continue
		}
		repoType := d.GetRepoInterfaceType()
		repoField := golang.NewField(d.GetExternalRepoName(), repoType, true)
		containerStruct.AddField(repoField)

		services := domainServices{
			repo: repoField,
		}

		if d.HasService() {
			serviceType := d.GetServiceInterfaceType()
			serviceField := golang.NewField(d.GetExternalServiceName(), serviceType, true)
			containerStruct.AddField(serviceField)
			services.service = serviceField
		}

		result.fields[resourceName] = services
	}
	result.AddVar(golang.NewVar(o.SingletonName, "", containerStruct.CopyType(), false))
	result.AddStruct(containerStruct)

	return result
}

func (c *Container) GetContainerType() *golang.Struct {
	return c.containerType
}

func (c *Container) GetConstructor() *golang.Function {
	return c.containerType.GetConstructor()
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
