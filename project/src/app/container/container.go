package container

import "github.com/68696c6c/capricorn_rnd/golang"

type Container struct {
	*golang.File
}

func NewContainer(pkg *golang.Package) Container {
	return Container{
		File: pkg.AddGoFile("app"),
	}
}
