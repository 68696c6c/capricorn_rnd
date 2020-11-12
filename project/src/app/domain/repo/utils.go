package repo

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type methodMeta struct {
	modelType           model.Type
	modelPlural         string
	modelArgName        string
	baseImport          string
	pkgName             string
	receiverName        string
	dbFieldRef          string
	dbFieldName         string
	dbType              golang.Type
	queryArgName        string
	queryType           golang.Type
	baseQueryFuncName   string
	filterQueryFuncName string
	pageQueryFuncName   string
	repoStructType      golang.Type
	repoInterfaceType   golang.Type
}

func makeMethodMeta(meta model.Meta, baseImport, pkgName, recName string, repoStruct, repoInterface golang.Type) methodMeta {
	dbFieldName := "db"
	return methodMeta{
		modelType:           determineModelType(pkgName, meta.ModelType),
		modelPlural:         strings.ToLower(utils.Plural(meta.ModelType.Name)),
		modelArgName:        "model",
		baseImport:          baseImport,
		pkgName:             pkgName,
		receiverName:        recName,
		dbFieldRef:          fmt.Sprintf("%s.%s", recName, dbFieldName),
		dbFieldName:         dbFieldName,
		dbType:              golang.MakeDbConnectionType(),
		queryArgName:        "query",
		queryType:           golang.MakeQueryType(),
		baseQueryFuncName:   "getBaseQuery",
		filterQueryFuncName: "getFilteredQuery",
		pageQueryFuncName:   "applyPaginationToQuery",
		repoStructType:      repoStruct,
		repoInterfaceType:   repoInterface,
	}
}

// Determine the correct reference to the model.
func determineModelType(repoPkgName string, rawModelType model.Type) model.Type {
	result := rawModelType

	// Repos always operate on pointers to models.
	result.IsPointer = true

	// If we are generating a DDD app, the model and repo will be in the same package and references to the model in the
	// repo file should not include the package name.
	if result.Package == repoPkgName {
		result.Package = ""
	}

	return result
}
