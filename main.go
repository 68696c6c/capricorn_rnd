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

	g := generator.NewGenerator(generator.PanicHandler{})
	g.Generate(module.Build("examples"))
}
