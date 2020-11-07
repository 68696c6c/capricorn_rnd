package utils

import "path"

// Folder represents an internal node in a project tree.
type Folder struct {
	FullPath string
	BasePath string
	Name     string
	folders  []Directory
	files    []RenderableFile
}

func NewFolder(basePath, name string) *Folder {
	return &Folder{
		FullPath: path.Join(basePath, name),
		BasePath: basePath,
		Name:     name,
		folders:  []Directory{},
		files:    []RenderableFile{},
	}
}

func (f *Folder) AddFolder(name string) *Folder {
	dir := NewFolder(f.BasePath, name)
	f.folders = append(f.folders, dir)
	return dir
}

func (f *Folder) AddFile(name, ext string) *File {
	file := NewFile(f.BasePath, name, ext)
	f.files = append(f.files, file)
	return file
}

func (f *Folder) GetPath() string {
	return f.FullPath
}

func (f *Folder) GetFiles() []RenderableFile {
	return f.files
}

func (f *Folder) GetDirectories() []Directory {
	return f.folders
}

func (f *Folder) AddDirectory(d Directory) {
	f.folders = append(f.folders, d)
}

func (f *Folder) AddRenderableFile(r RenderableFile) {
	f.files = append(f.files, r)
}
