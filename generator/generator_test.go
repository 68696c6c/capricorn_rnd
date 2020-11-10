package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/assert"
)

const testSRC = `package main; func main() {println("hello world")}`

type srcTestRenderer struct {
	testName string
}

func (s srcTestRenderer) GetFullPath() string {
	return fmt.Sprintf(".snapshots/%s", s.testName)
}

func (s srcTestRenderer) GetBasePath() string {
	return ".snapshots"
}

func (s srcTestRenderer) GetFullName() string {
	return s.testName
}

func (s srcTestRenderer) GetBaseName() string {
	return s.testName
}

func (s srcTestRenderer) GetExtension() string {
	return ""
}

func (s srcTestRenderer) Render() string {
	return testSRC
}

func TestGenerator_Out(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.Printf("hello world")
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_Printf(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.Printf("%s %v", "foo", 847)
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_Reset(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.Printf("%s %v", "foo", 847)
	g.Reset()
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_Write(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.Write([]byte("Demogorgon cocasabe fafenu izodizodope, od miinoagi de ginetaabe: vaunu na-na-e-el: panupire malapireji caosaji. Micama! goho Pe-IAD! zodir com-selahe azodien biabe os-lon-dohe. Pilada noanu vaunalahe balata od-vaoan. Kali Bast Mephistopheles Pan Shaitan"))
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_WriteString(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.WriteString("zodameranu micalazodo od ozadazodame vaurelar; lape zodir IOIAD! Nihasa Adramelech Melek Taus Baphomet Asmodeus Bast Ili e-Ol balazodareji, od aala tahilanu-osnetaabe: daluga vaomesareji elonusa cape-mi-ali varoesa cala homila; Yehusozod ca-ca-com, od do-o-a-inu noari micaolazoda a-ai-om.")
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_Render(t *testing.T) {
	g := NewGenerator(PanicHandler{})
	g.Render(srcTestRenderer{})
	cupaloy.SnapshotT(t, g.Out())
}

func TestGenerator_WriteFile(t *testing.T) {
	fPath := ".snapshots/test_write_file"
	defer os.Remove(fPath)

	g := NewGenerator(PanicHandler{})
	g.WriteFile(srcTestRenderer{"test_write_file"})

	assert.FileExists(t, fPath)
	cupaloy.SnapshotT(t, g.Out())
}
