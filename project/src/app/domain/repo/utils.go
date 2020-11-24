package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type methodMeta struct {
	modelType           golang.IType
	modelPlural         string
	modelArgName        string
	receiverName        string
	dbFieldRef          string
	dbFieldName         string
	dbType              *golang.Type
	queryArgName        string
	queryType           *golang.Type
	baseQueryFuncName   string
	filterQueryFuncName string
	filterFuncName      string
	pageQueryFuncName   string
	repoStructType      *golang.Struct
	repoInterfaceType   *golang.Interface
	constructor         *golang.Function
}

func makeMethodMeta(domainMeta *config.DomainMeta, repoStruct *golang.Struct, repoInterface *golang.Interface) *methodMeta {
	recName := repoStruct.GetReceiverName()
	dbFieldName := "db"
	return &methodMeta{
		modelType:           domainMeta.GetModelType(),
		modelPlural:         utils.Camel(domainMeta.NamePlural),
		modelArgName:        "model",
		receiverName:        recName,
		dbFieldRef:          fmt.Sprintf("%s.%s", recName, dbFieldName),
		dbFieldName:         dbFieldName,
		dbType:              goat.MakeTypeDbConnection(),
		queryArgName:        "query",
		queryType:           goat.MakeTypeQuery(),
		baseQueryFuncName:   "getBaseQuery",
		filterQueryFuncName: "getFilteredQuery",
		filterFuncName:      "Filter",
		pageQueryFuncName:   "ApplyPaginationToQuery",
		repoStructType:      repoStruct,
		repoInterfaceType:   repoInterface,
	}
}

func (m *methodMeta) AddConstructor(constructor *golang.Function) {
	m.constructor = constructor
}
