package golang

import (
	"strings"

	"github.com/68696c6c/gonad/utils"
)

// File represents a leaf node in a golang project tree.
type File struct {
	*utils.File
}

func NewFile(name string, paths Paths) *File {
	nanes := utils.NewInflection(name)
	return &File{
		File: utils.NewFile(paths.File, nanes.Snake, "go"),
	}
}

// func (f *File) GetFullPath() string {
// 	return f.FullPath
// }
//
// func (f *File) GetBasePath() string {
// 	return f.BasePath
// }
//
// func (f *File) GetFullName() string {
// 	return f.FullName
// }
//
// func (f *File) GetBaseName() string {
// 	return f.BaseName
// }
//
// func (f *File) GetExtension() string {
// 	return f.Ext
// }

func (f *File) Render() []byte {
	lines := []string{
		"// Code generated by \"gonad\"; DO NOT EDIT.\n",
		// f.pkg.declaration,
		// string(f.imports.Render()),
	}
	return []byte(strings.Join(lines, "\n"))
}
