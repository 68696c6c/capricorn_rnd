package cmd

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/utils"
)

type Map map[string]Command

type CMD struct {
	pkg      *golang.Package
	commands Map
}

type Command struct {
	file *golang.File
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

func NewCMD(pkg *golang.Package, commands []Command) CMD {
	result := make(Map)
	pkgCmd := pkg.AddPackage("cmd")
	for _, e := range commands {
		key := utils.Kebob(e.Name)
		result[key] = newCommand(pkgCmd, e)
	}
	return CMD{
		pkg:      pkgCmd,
		commands: result,
	}
}

func newCommand(pkg *golang.Package, c Command) Command {
	return Command{
		file: pkg.AddGoFile(c.Name),
		Name: c.Name,
		Args: c.Args,
	}
}
