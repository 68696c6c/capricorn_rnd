package utils

import "path"

// Folder represents an internal node in a project tree.
type Folder struct {
	fullPath string
	basePath string
	name     string
	folders  []Directory
	files    []RenderableFile
}

func NewFolder(basePath, name string) *Folder {
	return &Folder{
		fullPath: path.Join(basePath, name),
		basePath: basePath,
		name:     name,
		folders:  []Directory{},
		files:    []RenderableFile{},
	}
}

func (f *Folder) AddFolder(name string) *Folder {
	dir := NewFolder(f.basePath, name)
	f.folders = append(f.folders, dir)
	return dir
}

func (f *Folder) AddFile(name, ext string) *File {
	file := NewFile(f.basePath, name, ext)
	f.files = append(f.files, file)
	return file
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetFullPath() string {
	return f.fullPath
}

func (f *Folder) GetBasePath() string {
	return f.basePath
}

func (f *Folder) GetImport() string {
	return f.fullPath
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
