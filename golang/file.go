package golang

import (
	"fmt"
	"github.com/68696c6c/gonad/utils"
	"path"
	"strings"
)

// File represents a leaf node in a golang project tree.
type File struct {
	Paths
	// imports  *imports
	// pkg      *Package
	// Paths     string
	fullPath string
	basePath string
	fullName string
	baseName string
	ext      string
}

func NewFile(name, ext string, p Paths) *File {
	n := utils.NewInflection(name)
	baseName := n.Snake
	fullName := fmt.Sprintf("%s.%s", baseName, ext)
	return &File{
		fullName: fullName,
		baseName: baseName,
		ext:      ext,
		fullPath: path.Join(p.File, fullName),
		basePath: p.File,
	}
}

func (f *File) GetFullPath() string {
	return f.fullPath
}

func (f *File) GetBasePath() string {
	return f.basePath
}

func (f *File) GetFullName() string {
	return f.fullName
}

func (f *File) GetBaseName() string {
	return f.baseName
}

func (f *File) GetExtension() string {
	return f.ext
}

func (f *File) Render() []byte {
	lines := []string{
		"// Code generated by \"gonad\"; DO NOT EDIT.\n",
		// f.pkg.declaration,
		// string(f.imports.Render()),
	}
	return []byte(strings.Join(lines, "\n"))
}
