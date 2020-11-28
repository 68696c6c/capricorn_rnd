package cmd

import (
	"path/filepath"

	"github.com/68696c6c/capricorn_rnd/generator"
	"github.com/68696c6c/capricorn_rnd/project"
	"github.com/68696c6c/capricorn_rnd/project/config"

	"github.com/spf13/cobra"
)

var GenerateGoat = &cobra.Command{
	Use:   "goat [spec] [destination]",
	Short: "Generates a new Goat project",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.ExactArgs(2)

		specPath := args[0]
		destinationPath := args[1]

		p, err := config.NewProjectFromSpec(specPath)
		if err != nil {
			panic(err)
		}

		g := generator.NewGenerator(generator.PanicHandler{})
		projectDir := project.Build(p, config.NewProjectOptions(destinationPath))
		g.Generate(projectDir)

		projectPath := filepath.Join(projectDir.GetFullPath(), "src")
		err = project.FMT(projectPath)
		if err != nil {
			panic(err)
		}

		err = project.InitModule(projectPath, p.Module)
		if err != nil {
			panic(err)
		}

		err = project.Setup(projectDir.GetFullPath())
		if err != nil {
			panic(err)
		}
	},
}
