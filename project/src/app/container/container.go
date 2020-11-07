package container

import "github.com/68696c6c/gonad/golang"

type Container struct {
	file *golang.File
}

func NewContainer(pkg *golang.Package) Container {
	return Container{
		file: pkg.AddGoFile("app"),
	}
}
