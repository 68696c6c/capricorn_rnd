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
	repoStructType      *golang.Type
	repoInterfaceType   *golang.Type
}

func makeMethodMeta(domainMeta *config.DomainResource, recName string, repoStruct, repoInterface *golang.Type) methodMeta {
	dbFieldName := "db"
	return methodMeta{
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
