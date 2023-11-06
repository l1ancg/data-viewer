go-env:
	go env

build: 
	go build -ldflags="-s -w" -o data-viewer cmd/main.go
	pwd
	ls -hal | grep data-viewer

wire:
	$(shell go env GOPATH)/bin/wire cmd/wire/gen.go

go-run:
	go run cmd/main.go

