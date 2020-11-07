package golang

import (
	"fmt"
	"path"

	"github.com/68696c6c/capricorn_rnd/utils"
)

// Package represents an internal node in a golang project tree.
type Package struct {
	Paths
	packages    []utils.Directory
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

func (p *Package) AddPackage(name string) *Package {
	pkg := NewPackage(name, p.Paths)
	p.packages = append(p.packages, pkg)
	return pkg
}

func (p *Package) AddGoFile(name string) *File {
	file := NewFile(name, p.Paths)
	p.files = append(p.files, file)
	return file
}

func (p *Package) AddFile(name, ext string) *utils.File {
	file := utils.NewFile(p.GetPath(), name, ext)
	p.files = append(p.files, file)
	return file
}

func (p *Package) GetPath() string {
	return p.Paths.File
}

func (p *Package) GetFiles() []utils.RenderableFile {
	return p.files
}

func (p *Package) GetDirectories() []utils.Directory {
	return p.packages
}

func (p *Package) AddFolder(name string) *utils.Folder {
	folder := utils.NewFolder(p.GetPath(), name)
	p.packages = append(p.packages, folder)
	return folder
}
