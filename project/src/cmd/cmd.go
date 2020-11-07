package cmd

import (
	"github.com/68696c6c/capricorn_rnd/golang"
)

type CMD struct {
	*golang.Package
}

type Command struct {
	file *golang.File
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

func Build(pkg *golang.Package, commands []Command) {
	pkgCmd := pkg.AddPackage("cmd")
	for _, c := range commands {
		pkgCmd.AddGoFile(c.Name)
	}
}
