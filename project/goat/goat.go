package goat

import "github.com/68696c6c/capricorn_rnd/golang"

const (
	ImportGoat       = "github.com/68696c6c/goat"
	ImportQuery      = "github.com/68696c6c/goat/query"
	ImportGin        = "github.com/gin-gonic/gin"
	ImportErrors     = "github.com/pkg/errors"
	ImportGorm       = "github.com/jinzhu/gorm"
	ImportValidation = "github.com/go-ozzo/ozzo-validation"
	ImportLogrus     = "github.com/sirupsen/logrus"
	ImportGoose      = "github.com/pressly/goose"
	ImportCobra      = "github.com/spf13/cobra"
	ImportViper      = "github.com/spf13/viper"
	ImportSqlDriver  = "_ \"github.com/go-sql-driver/mysql\""
)

func MakeTypeId() *golang.Type {
	return golang.MockType(ImportGoat, "ID", false, false)
}

func MakeTypeQuery() *golang.Type {
	return golang.MockType(ImportQuery, "Query", true, false)
}

func MakeTypeDbConnection() *golang.Type {
	return golang.MockType(ImportGorm, "DB", true, false)
}

func MakeTypeLogger() *golang.Type {
	return golang.MockType(ImportLogrus, "Logger", true, false)
}

func MakeTypeErrorHandler() *golang.Type {
	return golang.MockType(ImportGoat, "ErrorHandler", false, false)
}

func MakeTypeHandlerFunc() *golang.Type {
	return golang.MockType(ImportGin, "HandlerFunc", false, false)
}

func MakeTypeGinContext() *golang.Type {
	return golang.MockType(ImportGin, "Context", true, false)
}

func MakeTypeRouter() *golang.Type {
	return golang.MockType(ImportGoat, "Router", false, false)
}

func MakeHardModelStruct() *golang.Struct {
	result := golang.StructFromType(golang.MockType(ImportGoat, "Model", false, false))
	result.AddField(MakeModelField("id", MakeTypeId(), true, false, true))
	result.AddField(MakeModelField("created_at", golang.MakeTypeTime(false), true, false, true))
	result.AddField(MakeModelField("updated_at", golang.MakeTypeTime(true), true, false, true))
	return result
}

func MakeSoftModelStruct() *golang.Struct {
	result := MakeHardModelStruct()
	result.AddField(MakeModelField("deleted_at", golang.MakeTypeTime(true), true, false, true))
	return result
}

func MakeModelField(separatedName string, t golang.IType, isExported, isRequired, omitEmpty bool) *golang.Field {
	result := golang.NewField(separatedName, t, isExported)
	result.SetRequired(isRequired)
	result.SetJsonTag(omitEmpty)
	return result
}
