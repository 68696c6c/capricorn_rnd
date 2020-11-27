package config

import "github.com/68696c6c/capricorn_rnd/golang"

type Command struct {
	*golang.File
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

type CmdOptions struct {
	PkgName string

	CmdArgName  string
	ArgsArgName string

	RootFileName   string
	RootVarName    string
	RootCommandUse string

	ServerFileName   string
	ServerCommandUse string

	MigrateFileName   string
	MigrateCommandUse string
}
