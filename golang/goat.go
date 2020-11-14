package golang

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

func makeBaseModelType() Type {
	return NewTypeMock(ImportGoat, "Model", false, false)
}

func MakeIdType() Type {
	return NewTypeMock(ImportGoat, "ID", false, false)
}

func MakeErrorType() Type {
	return NewTypeMock("", "error", false, false)
}

func MakeTimeType(isPointer bool) Type {
	return NewTypeMock("time", "Time", isPointer, false)
}

func MakeQueryType() Type {
	return NewTypeMock(ImportQuery, "Query", true, false)
}

func MakeDbConnectionType() Type {
	return NewTypeMock(ImportGorm, "DB", true, false)
}

func MakeHardModelStruct() *Struct {
	result := newStructFromType(makeBaseModelType())
	result.AddField(MakeModelField("id", MakeIdType(), true, false, true))
	result.AddField(MakeModelField("created_at", MakeTimeType(false), true, false, true))
	result.AddField(MakeModelField("updated_at", MakeTimeType(true), true, false, true))
	return result
}

func MakeSoftModelStruct() *Struct {
	result := MakeHardModelStruct()
	result.AddField(MakeModelField("deleted_at", MakeTimeType(true), true, false, true))
	return result
}

func MakeModelField(separatedName string, t IType, isExported, isRequired, omitEmpty bool) *Field {
	result := NewField(separatedName, t, isExported)
	result.SetRequired(isRequired)
	result.SetJsonTag(omitEmpty)
	return result
}
