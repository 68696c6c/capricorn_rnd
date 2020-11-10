package golang

import (
	"fmt"
	"path"

	"github.com/68696c6c/capricorn_rnd/utils"
)

// Package represents an internal node in a golang project tree.
type Package struct {
	Paths
	reference   string
	declaration string
	*utils.Folder
}

func NewPackage(name, basePath, baseImport string) *Package {
	pkgName := utils.Snake(name)
	return &Package{
		Folder: utils.NewFolder(basePath, pkgName),
		Paths: Paths{
			FilePath:   path.Join(basePath, pkgName),
			ImportPath: path.Join(baseImport, pkgName),
		},
		reference:   pkgName,
		declaration: fmt.Sprintf("package %s", pkgName),
	}
}

func (p *Package) AddPackage(name string) *Package {
	pkg := NewPackage(name, p.FilePath, p.ImportPath)
	p.AddDirectory(pkg)
	return pkg
}

func (p *Package) AddGoFile(name string) *File {
	file := NewFile(name, p.Paths)
	p.AddRenderableFile(file)
	file.PKG = p
	return file
}

func (p *Package) GetPath() string {
	return p.Paths.FilePath
}
