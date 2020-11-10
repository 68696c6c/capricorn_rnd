package golang

import (
	"fmt"
	"path"

	"github.com/68696c6c/capricorn_rnd/utils"
)

// Package represents an internal node in a golang project tree.
type Package struct {
	reference   string
	declaration string
	importPath  string
	baseImport  string
	*utils.Folder
}

func NewPackage(name, basePath, baseImport string) *Package {
	pkgName := utils.Snake(name)
	return &Package{
		Folder:      utils.NewFolder(basePath, pkgName),
		importPath:  path.Join(baseImport, pkgName),
		reference:   pkgName,
		declaration: fmt.Sprintf("package %s", pkgName),
		baseImport:  baseImport,
	}
}

func (p *Package) AddPackage(name string) *Package {
	pkg := NewPackage(name, p.GetFullPath(), p.importPath)
	p.AddDirectory(pkg)
	return pkg
}

func (p *Package) AddGoFile(name string) *File {
	file := NewFile(p.GetFullPath(), name)
	p.AddRenderableFile(file)
	file.PKG = p
	return file
}

func (p *Package) GetImport() string {
	return p.importPath
}

func (p *Package) GetBaseImport() string {
	return p.baseImport
}
