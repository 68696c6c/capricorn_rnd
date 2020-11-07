package golang

import (
	"fmt"
	"path"

	"github.com/68696c6c/gonad/utils"
)

// Package represents an internal node in a golang project tree.
type Package struct {
	Paths
	packages    []utils.Package
	files       []utils.RenderableFile
	reference   string
	declaration string
}

func newPackage(name string, paths Paths) *Package {
	return &Package{
		Paths:       paths,
		reference:   name,
		declaration: fmt.Sprintf("package %s", name),
	}
}

func NewRootPackage(basePath, baseImport string) *Package {
	return newPackage("main", Paths{
		File:   basePath,
		Import: baseImport,
	})
}

func NewPackage(name string, base Paths) *Package {
	snake := utils.Snake(name)
	return newPackage(snake, Paths{
		File:   path.Join(base.File, snake),
		Import: path.Join(base.Import, snake),
	})
}

func (n *Package) AddFile(name, ext string) *File {
	file := NewFile(name, ext, n.Paths)
	n.files = append(n.files, file)
	return file
}

func (n *Package) AddPackage(name string) *Package {
	pkg := NewPackage(name, n.Paths)
	n.packages = append(n.packages, pkg)
	return pkg
}

func (n *Package) GetPath() string {
	return n.File
}

func (n *Package) GetFiles() []utils.RenderableFile {
	return n.files
}

func (n *Package) GetPackages() []utils.Package {
	return n.packages
}
