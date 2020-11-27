package container

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
)

type Config struct {
	*golang.File
}

func NewConfig(pkg golang.IPackage, o config.ServiceContainerConfigOptions) Config {
	return Config{
		File: pkg.AddGoFile(o.FileName),
	}
}
