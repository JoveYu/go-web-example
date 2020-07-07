GITTAG=$(shell git describe --abbrev=0 --tags)
GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse --short HEAD)

REGISTRY=test.com/example
IMAGE=${REGISTRY}:${GITBRANCH}-${GITCOMMIT}

swagger:
	command -v swag >/dev/null 2>&1  || go get -u github.com/swaggo/swag/cmd/swag 
	swag init

build: swagger
	go build -o main

run: swagger
	go run main.go

watch:
	command -v CompileDaemon >/dev/null 2>&1  || go get -u github.com/githubnemo/CompileDaemon
	CompileDaemon -build="make build" -command=./main -exclude-dir=.git -exclude-dir=docs

docker.build:
	docker build -t ${IMAGE} .

docker.run:
	docker run --rm -it -v ${PWD}/config.yaml:/app/config.yaml ${IMAGE}

docker.push:
	docker push ${IMAGE}
