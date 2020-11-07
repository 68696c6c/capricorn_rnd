package container

import "github.com/68696c6c/capricorn_rnd/golang"

type Config struct {
	file *golang.File
}

func NewConfig(pkg *golang.Package) Config {
	return Config{
		file: pkg.AddGoFile("config"),
	}
}
