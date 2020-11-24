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
}

func Build(pkg golang.IPackage, fileName string, domainMeta *config.DomainMeta) *Service {
	actions := domainMeta.GetServiceActions()
	if len(actions) == 0 {
		return nil
	}

	baseTypeName := utils.Pascal(fileName)
	serviceStruct := golang.NewStruct(baseTypeName+"Implementation", false, false)
	serviceInterface := golang.NewInterface(baseTypeName, false, false)

	repoFieldName := "repo"
	repoType := domainMeta.GetRepoType()
	serviceStruct.AddConstructor(makeConstructor(serviceStruct.Type, serviceInterface.Type, repoType, repoFieldName))

	result := &Service{
		File:          pkg.AddGoFile(fileName),
		structType:    serviceStruct,
		interfaceType: serviceInterface,
	}
	result.AddStruct(serviceStruct)
	result.AddInterface(serviceInterface)

	serviceStruct.AddField(golang.NewField(repoFieldName, repoType, false))

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
