package golang

import (
	"fmt"
	"path/filepath"
	"strings"
)

const DefaultPackageString = "???"

type IType interface {
	GetType() *Type
	CopyType() *Type
	GetImport() string
	GetReference() string
	GetName() string
	GetPackage() string
	SetPackage(pkgName string)
	GetIsPointer() bool
	GetIsSlice() bool
	GetStructFields() Fields
	SetReceiver(v *Value)
	GetReceiverName() string
	AddFunction(f *Function)
	getImports() imports
}

type Type struct {
	Import    string
	Package   string
	Name      string
	IsPointer bool
	IsSlice   bool
	*imports
	receiver  *Value
	functions Functions
}

func (t *Type) GetType() *Type {
	return t
}

func (t *Type) CopyType() *Type {
	return copyType(t)
}

func (t *Type) GetImport() string {
	return t.Import
}

func (t *Type) getImports() imports {
	return *t.imports
}

func (t *Type) GetReference() string {
	prefix := ""
	if t.IsSlice {
		prefix += "[]"
	}
	if t.IsPointer {
		prefix += "*"
	}
	if t.Package == "" {
		return prefix + t.Name
	}
	return fmt.Sprintf("%s%s.%s", prefix, t.Package, t.Name)
}

func (t *Type) GetName() string {
	return t.Name
}

func (t *Type) GetPackage() string {
	return t.Package
}

func (t *Type) SetPackage(pkgName string) {
	t.Package = pkgName
}

func (t *Type) GetIsPointer() bool {
	return t.IsPointer
}

func (t *Type) GetIsSlice() bool {
	return t.IsSlice
}

func (t *Type) GetStructFields() Fields {
	return Fields{}
}

// Sets the receiver to a copy of this type, minus the Import and Package since a receiver will never need those.
func (t *Type) initReceiver() {
	imps := t.getImports()
	name := t.GetName()
	recName := "r"
	if len(name) > 0 {
		recName = strings.ToLower(name[0:1])
	}
	t.receiver = ValueFromType(recName, &Type{
		Import:    "",
		Package:   "",
		Name:      name,
		IsPointer: t.GetIsPointer(),
		IsSlice:   t.GetIsSlice(),
		imports:   &imps,
	})
}

func (t *Type) SetReceiver(v *Value) {
	t.receiver = v
}

func (t *Type) GetReceiverName() string {
	return t.receiver.Name
}

func (t *Type) AddFunction(f *Function) {
	f.SetReceiver(t.receiver)
	t.imports = mergeImports(*t.imports, f.getImports())
	t.functions = append(t.functions, f)
}

func copyType(t *Type) *Type {
	imps := t.getImports()
	result := &Type{
		Import:    t.GetImport(),
		Package:   t.GetPackage(),
		Name:      t.GetName(),
		IsPointer: t.GetIsPointer(),
		IsSlice:   t.GetIsSlice(),
		imports:   &imps,
	}
	result.initReceiver()
	return result
}

// Use this for generating types.  The import and package will be set when the Type is added to a File and should never be referenced before that happens..
func NewType(typeName string, isPointer, isSlice bool) *Type {
	result := &Type{
		Import:    DefaultPackageString,
		Package:   DefaultPackageString,
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
		imports:   newImports(),
	}
	result.initReceiver()
	return result
}

// Use this function for mocking built in or vendor types.
func MockType(importPath, typeName string, isPointer, isSlice bool) *Type {
	pkgName := ""
	if importPath != "" {
		pkgName = filepath.Base(importPath)
	}
	result := &Type{
		Import:    importPath,
		Package:   pkgName,
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
		imports:   newImports(),
	}
	result.initReceiver()
	return result
}

// AVOID USING THIS IF POSSIBLE
// @TODO This is currently only used for generating user-defined model fields; setting the import correctly will require defining a known set of supported types.
func MockTypeFromReference(reference string) IType {
	trimmed, isSlice, isPointer := isReferenceSliceOrPointerAndTrim(reference)
	pkgName, typeName := getPkgAndTypeFromReference(trimmed)
	result := &Type{
		Import:    "???",
		Package:   pkgName,
		Name:      typeName,
		IsPointer: isPointer,
		IsSlice:   isSlice,
		imports:   newImports(),
	}
	result.initReceiver()
	return result
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
