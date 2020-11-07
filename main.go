package main

import (
	"fmt"

	"github.com/68696c6c/gonad/generator"
	"github.com/68696c6c/gonad/project"
)

func main() {
	var err error
	exampleName := "basic"
	module, err := project.NewProjectFromSpec(fmt.Sprintf("%s.yml", exampleName))
	if err != nil {
		panic(err)
	}

	examplePath := fmt.Sprintf("examples/%s", exampleName)
	g := generator.NewGenerator(generator.PanicHandler{})
	g.Generate(module.Build(examplePath))

	// buf := &bytes.Buffer{}
	// memviz.Map(buf, &module.Project)
	// dataPath := fmt.Sprintf("%s/%s", examplePath, "diagram-data")
	// err = ioutil.WriteFile(dataPath, buf.Bytes(), 0644)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// diagramPath := fmt.Sprintf("%s/%s", examplePath, "diagram.png")
	// cmd := exec.Command("dot", "-Tpng", dataPath, "-o", diagramPath)
	// err = cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }
}
