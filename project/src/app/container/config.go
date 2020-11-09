package container

import "github.com/68696c6c/capricorn_rnd/golang"

type Config struct {
	*golang.File
}

func NewConfig(pkg *golang.Package) Config {
	return Config{
		File: pkg.AddGoFile("config"),
	}
}
