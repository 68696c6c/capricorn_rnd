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

  app:
    image: app:dev
    command: ./app server 80
    depends_on:
      - db
    volumes:
      - pkg:/go/pkg
      - ./:/{{ .Workdir }}
    working_dir: /{{ .Workdir }}
    ports:
      - "80"
    env_file:
      - .app.env
    environment:
      VIRTUAL_HOST: {{ .AppHTTPAlias }}.local
      ENV: local
      HTTP_PORT: 80
      MIGRATION_PATH: /{{ .Workdir }}/src/database
    networks:
      default:
        aliases:
          - {{ .AppHTTPAlias }}.local

  db:
    image: mysql:5.7
    environment:
      MYSQL_ROOT_PASSWORD: {{ .MainDatabase.Password }}
      MYSQL_DATABASE: {{ .MainDatabase.Name }}
    ports:
      - "${HOST_DB_PORT:-3310}:{{ .MainDatabase.Port }}"
    volumes:
      - db-volume:/var/lib/mysql
`

type DockerCompose struct {
	*utils.File
	data config.Ops
}

func NewDockerCompose(basePath string, c config.Ops) DockerCompose {
	file := utils.NewFile(basePath, "docker-compose", "yml")
	return DockerCompose{
		File: file,
		data: c,
	}
}

func (d DockerCompose) Render() string {
	result, err := utils.ParseTemplate(d.FullPath, dockerComposeTemplate, d.data)
	if err != nil {
		panic(err)
	}
	return result
}
