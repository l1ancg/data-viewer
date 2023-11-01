go-env:
	go env

build: 
	go build -ldflags="-s -w" -o data-viewer backend/cmd/main.go
	pwd
	ls -hal | grep data-viewer

wire:
	$(shell go env GOPATH)/bin/wire backend/cmd/wire/gen.go

go-run:
	go run backend/cmd/main.go

