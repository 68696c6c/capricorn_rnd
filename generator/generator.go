package generator

import (
	"bytes"
	"fmt"

	"github.com/68696c6c/capricorn_rnd/utils"
)

func NewGenerator(e ErrorHandler) utils.Generator {
	return &generator{
		buffer: &bytes.Buffer{},
		errors: e,
	}
}

type generator struct {
	buffer *bytes.Buffer
	errors ErrorHandler
}

// Writes the provided format string to the buffer.
func (g *generator) Printf(format string, args ...interface{}) utils.Generator {
	_, err := fmt.Fprintf(g.buffer, format, args...)
	g.errors.HandleError(err)
	return g
}

// Clears the buffer.
func (g *generator) Reset() utils.Generator {
	g.buffer.Reset()
	return g
}

func (g *generator) Write(b []byte) utils.Generator {
	g.buffer.Write(b)
	return g
}

func (g *generator) WriteString(s string) utils.Generator {
	g.buffer.Write([]byte(s))
	return g
}

// Appends the provided Renderable to the output.
func (g *generator) Render(r utils.Renderable) utils.Generator {
	g.Write(r.Render())
	return g
}

// Returns the current buffer contents.
func (g *generator) Out() []byte {
	return g.buffer.Bytes()
}

// Sets and writes the provided RenderableFile to it's target file.
func (g *generator) WriteFile(r utils.RenderableFile) []byte {
	out := r.Render()
	result := g.Reset().Write(out).Out()
	err := writeFile(r.GetFullPath(), result)
	g.errors.HandleError(err)
	return result
}

func (g *generator) Generate(d utils.Directory) {
	err := createDir(d.GetPath())
	if err != nil {
		panic(err)
	}
	for _, file := range d.GetFiles() {
		g.WriteFile(file)
	}
	for _, dir := range d.GetDirectories() {
		g.Generate(dir)
	}
}
