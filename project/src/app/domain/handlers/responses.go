package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func makeResourceResponse(name string, modelType *golang.Struct) *golang.Struct {
	result := golang.NewStruct(name, false, false)
	result.AddField(golang.NewField("", modelType, true))
	return result
}

func makeListResponse(name string, modelType *golang.Struct) *golang.Struct {
	result := golang.NewStruct(name, false, false)

	mt := modelType.CopyType()
	mt.IsPointer = true
	sliceType := golang.MakeSliceType(false, mt)
	dataField := golang.NewField("Data", sliceType, true)
	dataField.AddTag("json", []string{"data"})
	result.AddField(dataField)

	queryField := golang.NewField("", goat.MakeTypePagination(), true)
	queryField.AddTag("json", []string{"pagination"})
	result.AddField(queryField)

	return result
}
