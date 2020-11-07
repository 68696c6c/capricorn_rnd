package local

import "github.com/68696c6c/capricorn_rnd/utils"

// @TODO use a hosted base image
const dockerfileTemplate = `FROM golang:1.15-alpine as env

ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org,direct

RUN apk add --no-cache git gcc openssh mysql-client

RUN go get golang.org/x/tools/cmd/goimports

# Install swagger for generating API docs.
RUN go get -v github.com/go-swagger/go-swagger/cmd/swagger

# Install golangci-lint for linting.
RUN wget https://github.com/golangci/golangci-lint/releases/download/v1.24.0/golangci-lint-1.24.0-linux-amd64.tar.gz \
    && tar xzf golangci-lint-1.24.0-linux-amd64.tar.gz \
    && mv golangci-lint-1.24.0-linux-amd64/golangci-lint /usr/local/bin/golangci-lint

# Install goose for running migrations.
RUN go get -u github.com/pressly/goose/cmd/goose

RUN mkdir -p /{{ .Workdir }}
WORKDIR /{{ .Workdir }}


# Local development stage.
FROM env as dev

RUN go get -v github.com/go-delve/delve/cmd/dlv

RUN apk add --no-cache bash
RUN echo 'alias ll="ls -lah"' >> ~/.bashrc


# Pipeline stage for running unit tests, linters, etc.
FROM env as built

COPY ./src /{{ .Workdir }}
RUN go build -i -o app
`

type Dockerfile struct {
	*utils.File
	data Config
}

func NewDockerfile(basePath string, c Config) Dockerfile {
	file := utils.NewFile(basePath, "Dockerfile", "")
	return Dockerfile{
		File: file,
		data: c,
	}
}

func (d Dockerfile) Render() []byte {
	result, err := utils.ParseTemplate(d.FullPath, dockerfileTemplate, d.data)
	if err != nil {
		panic(err)
	}
	return result
}
