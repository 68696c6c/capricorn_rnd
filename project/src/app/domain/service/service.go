package service

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Meta struct {
	Methods  []string
	RepoType *golang.Type
}

type Service struct {
	*golang.File
	structType    *golang.Struct
	interfaceType *golang.Interface
	externalName  string
}

func Build(pkg golang.IPackage, o config.ServiceOptions, domainMeta *config.DomainMeta) *Service {
	actions := domainMeta.GetServiceActions()
	if len(actions) == 0 {
		return nil
	}

	intTypeName := utils.Pascal(o.InterfaceNameTemplate.Parse(domainMeta.ResourceName))
	serviceInterface := golang.NewInterface(intTypeName, false, false)

	impTypeName := utils.Pascal(o.ImplementationNameTemplate.Parse(domainMeta.ResourceName))
	serviceStruct := golang.NewStruct(impTypeName, false, false)

	repoType := domainMeta.GetRepoType()
	serviceStruct.AddConstructor(makeConstructor(serviceStruct.Type, serviceInterface.Type, repoType, o.RepoFieldName))

	fileName := o.FileNameTemplate.Parse(domainMeta.ResourceName)
	result := &Service{
		File:          pkg.AddGoFile(fileName),
		structType:    serviceStruct,
		interfaceType: serviceInterface,
		externalName:  o.ExternalNameTemplate.Parse(domainMeta.ResourceName),
	}
	result.AddStruct(serviceStruct)
	result.AddInterface(serviceInterface)
	result.AddImportsApp(domainMeta.ImportRepos)

	serviceStruct.AddField(golang.NewField(o.RepoFieldName, repoType, false))

	for _, c := range actions {
		m := golang.NewFunction(utils.Pascal(c))
		serviceStruct.AddFunction(m)
		serviceInterface.AddFunction(m)
	}

	return result
}

func (s *Service) GetInterfaceType() *golang.Interface {
	return s.interfaceType
}

func (s *Service) GetConstructor() *golang.Function {
	return s.structType.GetConstructor()
}

func (s *Service) GetExternalName() string {
	return s.externalName
}
