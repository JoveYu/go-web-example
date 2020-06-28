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

