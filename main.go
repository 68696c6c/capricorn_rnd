package main

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/generator"
	"github.com/68696c6c/capricorn_rnd/project"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

func main() {
	exampleName := "complex"
	module, err := config.NewProjectFromSpec(fmt.Sprintf("%s.yml", exampleName))
	if err != nil {
		panic(err)
	}

	g := generator.NewGenerator(generator.PanicHandler{})
	g.Generate(project.Build(module, config.NewProjectOptions("_examples")))
}
