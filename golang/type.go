package golang

import (
	"fmt"
	"strings"
)

type IType interface {
	GetImport() string
	GetReference() string
	GetName() string
	GetPackage() string
	GetIsPointer() bool
	GetIsSlice() bool
	GetStructFields() []Field
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

func (t Type) GetReference() string {
	if t.Package == "" {
		return t.Name
	}
	return fmt.Sprintf("%s.%s", t.Package, t.Name)
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

func (t Type) GetStructFields() []Field {
	return []Field{}
}

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
