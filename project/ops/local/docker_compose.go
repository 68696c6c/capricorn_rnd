package local

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const dockerComposeTemplate = `version: "3.6"

networks:
  default:
    external:
      name: docker-dev

volumes:
  pkg:
  db-volume:

services:

  {{ .ServiceNameApp }}:
    image: app:local
    command: ./{{ .AppBinaryName }} server
    depends_on:
      - {{ .ServiceNameDb }}
    volumes:
      - pkg:/go/pkg
      - .:/{{ .Config.Workdir }}
    working_dir: /{{ .Config.Workdir }}/src
    ports:
      - "80"
    env_file:
      - .app.env
    environment:
      VIRTUAL_HOST: {{ .Config.AppHTTPAlias }}
      ENV: local
      HTTP_PORT: 80
      MIGRATION_PATH: /{{ .Config.Workdir }}/src/db
    networks:
      default:
        aliases:
          - {{ .Config.AppHTTPAlias }}

  {{ .ServiceNameDb }}:
    image: mysql:5.7
    environment:
      MYSQL_ROOT_PASSWORD: {{ .Config.MainDatabase.Password }}
      MYSQL_DATABASE: {{ .Config.MainDatabase.Name }}
    ports:
      - "${HOST_DB_PORT:-3310}:{{ .Config.MainDatabase.Port }}"
    volumes:
      - db-volume:/var/lib/mysql
`

type DockerCompose struct {
	*utils.File
	data config.Ops
	meta config.OpsMeta
}

func NewDockerCompose(basePath string, c config.Ops, meta config.OpsMeta) DockerCompose {
	file := utils.NewFile(basePath, "docker-compose", "yml")
	return DockerCompose{
		File: file,
		data: c,
		meta: meta,
	}
}

func (d DockerCompose) Render() string {
	result, err := utils.ParseTemplate(d.FullPath, dockerComposeTemplate, struct {
		Config         config.Ops
		AppBinaryName  string
		ServiceNameApp string
		ServiceNameDb  string
	}{
		Config:         d.data,
		AppBinaryName:  d.meta.AppBinaryName,
		ServiceNameApp: d.meta.ServiceNameApp,
		ServiceNameDb:  d.meta.ServiceNameDb,
	})
	if err != nil {
		panic(err)
	}
	return result
}
