package local

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const makefileTemplate = `DCR = docker-compose run --rm

NETWORK_NAME ?= docker-dev
APP_NAME = {{ .ServiceNameApp }}
DB_NAME = {{ .ServiceNameDb }}

DOC_PATH_BASE = docs/swagger.json
DOC_PATH_FINAL = docs/api-spec.json

.PHONY: docs build

.DEFAULT:
	@echo 'App targets:'
	@echo
	@echo '    image-local        build the $(APP_NAME):local Docker image for local development'
	@echo '    image-built        build the $(APP_NAME):built Docker image for task running'
	@echo '    build              compile the app for use in Docker'
	@echo '    init               initialize the Go module'
	@echo '    deps               install dependencies'
	@echo '    setup-network      create local Docker network'
	@echo '    setup              set up local databases'
	@echo '    local              spin up local dev environment'
	@echo '    local-down         tear down local dev environment'
	@echo '    migrate            migrate the local database'
	@echo '    migration          create a new migration'
	@echo '    generate           generate String methods for app enum types'
	@echo '    docs               build the Swagger docs'
	@echo '    docs-server        build and serve the Swagger docs'
	@echo '    test               run unit tests'
	@echo '    lint               run the linter'
	@echo '    lint-fix           run the linter and fix any problems'
	@echo


default: .DEFAULT

image-local:
	docker build . --target dev -t $(APP_NAME):local

image-built:
	docker build . --target built -t $(APP_NAME):built

build:
	$(DCR) $(APP_NAME) go build -i -o {{ .AppBinaryName }}

deps:
	$(DCR) $(APP_NAME) go mod tidy
	$(DCR) $(APP_NAME) go mod vendor

setup-network:
	docker network create docker-dev || exit 0

setup: setup-network image-local generate deps build
	@test -f ".app.env" || (echo "you need to set up your .app.env file before running this command"; exit 1)
	$(DCR) $(DB_NAME) mysql -u root -p{{ .MainDatabasePassword }} -h $(DB_NAME) -e "CREATE DATABASE IF NOT EXISTS {{ .MainDatabaseName }}"
	$(DCR) $(APP_NAME) bash -c "./{{ .AppBinaryName }} migrate up"

local: local-down build
	docker-compose up $(APP_NAME)

local-down:
	docker-compose rm -sf

test:
	$(DCR) $(APP_NAME) go test ./... -cover

migrate: build
	$(DCR) $(APP_NAME) ./{{ .AppBinaryName }} migrate up

migration: build
	$(DCR) $(APP_NAME) goose -dir {{ .MigrationsPath }} create $(name)

generate:
	$(DCR) $(APP_NAME) go generate {{ .EnumsPath }}

docs: build
	$(DCR) $(APP_NAME) bash -c "GO111MODULE=off swagger generate spec -mo '$(DOC_PATH_BASE)'"
	$(DCR) $(APP_NAME) ./{{ .AppBinaryName }} gen:docs

docs-server: docs
	swagger serve "$(DOC_PATH_FINAL)"

lint:
	$(DCR) $(APP_NAME) golangci-lint run

lint-fix:
	$(DCR) $(APP_NAME) golangci-lint run --fix
`

type Makefile struct {
	*utils.File
	data config.Ops
	meta config.OpsMeta
}

func NewMakefile(basePath string, c config.Ops, meta config.OpsMeta) Makefile {
	file := utils.NewFile(basePath, "Makefile", "")
	return Makefile{
		File: file,
		data: c,
		meta: meta,
	}
}

func (m Makefile) Render() string {
	result, err := utils.ParseTemplate(m.FullPath, makefileTemplate, struct {
		MainDatabasePassword string
		MainDatabaseName     string
		MigrationsPath       string
		EnumsPath            string
		AppBinaryName        string
		ServiceNameApp       string
		ServiceNameDb        string
	}{
		MainDatabasePassword: m.data.MainDatabase.Password,
		MainDatabaseName:     m.data.MainDatabase.Name,
		MigrationsPath:       m.meta.ImportMigrations,
		EnumsPath:            m.meta.ImportEnums,
		AppBinaryName:        m.meta.AppBinaryName,
		ServiceNameApp:       m.meta.ServiceNameApp,
		ServiceNameDb:        m.meta.ServiceNameDb,
	})
	if err != nil {
		panic(err)
	}
	return result
}
