package utils

import (
	"fmt"
	"path"
	"strings"
)

// File represents a leaf node in a project tree.
type File struct {
	FullPath string
	BasePath string
	FullName string
	BaseName string
	Ext      string
}

func NewFile(basePath, baseName, ext string) *File {
	fullName := strings.TrimRight(fmt.Sprintf("%s.%s", baseName, ext), ".")
	return &File{
		FullName: fullName,
		BaseName: baseName,
		Ext:      ext,
		FullPath: path.Join(basePath, fullName),
		BasePath: basePath,
	}
}

func (f *File) GetFullPath() string {
	return f.FullPath
}

func (f *File) GetBasePath() string {
	return f.BasePath
}

func (f *File) GetFullName() string {
	return f.FullName
}

func (f *File) GetBaseName() string {
	return f.BaseName
}

func (f *File) GetExtension() string {
	return f.Ext
}

func (f *File) Render() string {
	return ""
}
