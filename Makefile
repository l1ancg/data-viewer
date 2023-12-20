go-env:
	go env

build: 
	go build -ldflags="-s -w" -o data-viewer cmd/main.go
	pwd
	ls -hal | grep data-viewer

wire:
	$(shell go env GOPATH)/bin/wire backend/cmd/wire/gen.go

backend-run:
	cd backend
	go run backend/cmd/main.go

backend-build:
	go build -ldflags="-s -w" -o data-viewer backend/cmd/main.go

frontend-run:
	cd frontend && pnpm install && pnpm run dev

frontend-build:
	cd frontend && pnpm install && pnpm run build




