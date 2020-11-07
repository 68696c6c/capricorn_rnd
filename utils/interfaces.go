package utils

type Generator interface {
	Printf(format string, args ...interface{}) Generator
	Reset() Generator
	Write(b []byte) Generator
	WriteString(s string) Generator
	Render(r Renderable) Generator
	Out() []byte
	WriteFile(r RenderableFile) []byte
	Generate(d Directory)
}

type Renderable interface {
	Render() []byte
}

type RenderableFile interface {
	Renderable

	// Returns the full path to the file, including the file name and extension.
	GetFullPath() string

	// Returns the path to the directory containing the file.
	GetBasePath() string

	// Returns file name, including extension.
	GetFullName() string

	// Returns file name without the extension.
	GetBaseName() string

	// Returns the file extension.
	GetExtension() string
}

type Directory interface {
	GetPath() string
	GetFiles() []RenderableFile
	GetDirectories() []Directory
	AddFile(name, ext string) *File
	AddFolder(name string) *Folder
	AddDirectory(d Directory)
	AddRenderableFile(r RenderableFile)
}
