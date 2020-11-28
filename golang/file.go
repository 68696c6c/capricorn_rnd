package golang

import (
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

// File represents a leaf node in a golang project tree.
type File struct {
	*utils.File
	*imports
	PKG *Package // this is set when a File is passed to package.AddGoFile()
	// InitFunction Function    `yaml:"init_function,omitempty"`
	// Consts       []Const     `yaml:"consts,omitempty"`
	vars []*Var
	// TypeAliases  []Value     `yaml:"type_aliases,omitempty"`
	iotas      []*Iota
	structs    []*Struct
	interfaces []*Interface
	functions  Functions
}

func NewFile(basePath, name string) *File {
	return &File{
		File:    utils.NewFile(basePath, utils.Snake(name), "go"),
		imports: newImports(),
	}
}

func (f *File) Render() string {
	var lines []string

	for _, v := range f.vars {
		f.imports = mergeImports(*f.imports, v.getImports())
		lines = append(lines, v.Render())
	}

	for _, i := range f.iotas {
		f.imports = mergeImports(*f.imports, i.getImports())
		lines = append(lines, i.Render())
	}

	for _, i := range f.interfaces {
		f.imports = mergeImports(*f.imports, i.getImports())
		lines = append(lines, i.Render())
	}

	for _, s := range f.structs {
		f.imports = mergeImports(*f.imports, s.getImports())
		lines = append(lines, s.Render())
	}

	functionLines := f.functions.Render()
	if functionLines != "" {
		lines = append(lines, "")
		lines = append(lines, functionLines)
	}

	result := []string{
		`// Code generated by "capricorn"; DO NOT EDIT.`,
		f.PKG.declaration,
		"",
		f.imports.Render(),
	}
	result = append(result, lines...)
	return strings.Join(result, "\n") + "\n"
}

func (f *File) AddVar(v *Var) {
	pkg := f.PKG.GetName()
	removePackageRefIfSamePackage(pkg, v)
	removePackageRefIfSamePackage(pkg, v.valueType)
	v.Type.Package = pkg
	v.Type.Import = f.PKG.GetImport()
	f.vars = append(f.vars, v)
}

func (f *File) AddFunction(function *Function) {
	setFunctionPackages(f.PKG.GetName(), f.PKG.GetImport(), Functions{function})
	f.imports = mergeImports(*f.imports, function.getImports())
	f.functions = append(f.functions, function)
}

func (f *File) AddIota(i *Iota) {
	i.Type.Package = f.PKG.GetName()
	i.Type.Import = f.PKG.GetImport()
	f.iotas = append(f.iotas, i)
}

func (f *File) AddStruct(s *Struct) {
	pkg := f.PKG.GetName()
	imp := f.PKG.GetImport()
	s.Type.Package = pkg
	s.Type.Import = imp
	setFunctionPackages(pkg, imp, s.functions)
	for _, field := range s.GetStructFields() {
		removePackageRefIfSamePackage(pkg, field.Type)
	}
	f.structs = append(f.structs, s)
}

func (f *File) AddInterface(i *Interface) {
	pkg := f.PKG.GetName()
	imp := f.PKG.GetImport()
	i.Type.Package = pkg
	i.Type.Import = imp
	setFunctionPackages(pkg, imp, i.functions)
	f.interfaces = append(f.interfaces, i)
}

func setFunctionPackages(pkgName, importPath string, functions Functions) {
	for _, function := range functions {
		function.Type.Package = pkgName
		function.Type.Import = importPath
		for _, r := range function.returns {
			removePackageRefIfSamePackage(pkgName, r)
		}
		for _, a := range function.arguments {
			removePackageRefIfSamePackage(pkgName, a)
		}
	}
}

func removePackageRefIfSamePackage(filePkg string, subject IType) {
	subjectPkg := subject.GetPackage()
	if subjectPkg == filePkg || subjectPkg == DefaultPackageString {
		subject.SetPackage("")
	}
}
