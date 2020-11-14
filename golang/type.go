package golang

import (
	"fmt"
	"path/filepath"
	"strings"
)

type IType interface {
	GetImport() string
	GetReference() string
	GetName() string
	GetPackage() string
	GetIsPointer() bool
	GetIsSlice() bool
	GetStructFields() Fields
	GetImports() imports
}

type Type struct {
	Import    string
	Package   string
	Name      string
	IsPointer bool
	IsSlice   bool
}

func (t Type) GetImport() string {
	return t.Import
}

func (t Type) GetImports() imports {
	return imports{}
}

func (t Type) GetReference() string {
	prefix := ""
	if t.IsPointer {
		prefix += "*"
	}
	if t.IsSlice {
		prefix += "[]"
	}
	if t.Package == "" {
		return prefix + t.Name
	}
	return fmt.Sprintf("%s%s.%s", prefix, t.Package, t.Name)
}

func (t Type) GetName() string {
	return t.Name
}

func (t Type) GetPackage() string {
	return t.Package
}

func (t Type) GetIsPointer() bool {
	return t.IsPointer
}

func (t Type) GetIsSlice() bool {
	return t.IsSlice
}

func (t Type) GetStructFields() Fields {
	return Fields{}
}

// Use this for generating types.  The import and package will be set when the Type is added to a File and should never be referenced before that happens..
func NewType(typeName string, isPointer, isSlice bool) Type {
	return Type{
		Import:    "add this Type to a golang.File",
		Package:   "add this Type to a golang.File",
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
	}
}

// Use this function for mocking built in or vendor types.
func NewTypeMock(importPath, typeName string, isPointer, isSlice bool) Type {
	pkgName := ""
	if importPath != "" {
		pkgName = filepath.Base(importPath)
	}
	return Type{
		Import:    importPath,
		Package:   pkgName,
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
	}
}

// AVOID USING THIS IF POSSIBLE
// @TODO This is currently only used for generating user-defined model fields; setting the import correctly will require defining a known set of supported types.
func NewTypeFromReference(reference string) IType {
	trimmed, isSlice, isPointer := isReferenceSliceOrPointerAndTrim(reference)
	pkgName, typeName := getPkgAndTypeFromReference(trimmed)
	return Type{
		Import:    "???",
		Package:   pkgName,
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
	}
}

// Returns the provided reference string with any pointer or slice prefixes removed.
// Also returns boolean values indicating whether the reference was determined to be a pointer or slice.
// This function checks for both pointer and slice references because the checks for pointers and slices need to be done
// in the correct order.  I.e., the [] needs to be trimmed before we can check if the string starts with *.
func isReferenceSliceOrPointerAndTrim(reference string) (trimmedReference string, isSlice, isPointer bool) {
	result := reference
	if strings.HasPrefix(result, "[]") {
		isSlice = true
		result = strings.TrimPrefix(result, "[]")
	}
	if strings.HasPrefix(result, "*") {
		isPointer = true
		result = strings.TrimPrefix(result, "*")
	}
	return result, isSlice, isPointer
}

func getPkgAndTypeFromReference(trimmedReference string) (pkgName, typeName string) {
	if strings.Contains(trimmedReference, ".") {
		parts := strings.Split(trimmedReference, ".")
		return parts[0], parts[1]
	}
	return "", trimmedReference
}

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
