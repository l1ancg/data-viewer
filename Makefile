wire:
	$(shell go env GOPATH)/bin/wire backend/cmd/wire/gen.go

backend-run:
	cd backend
	go run backend/cmd/main.go

frontend-run:
	cd frontend && pnpm install && pnpm run dev

GO_BUILD = go build -ldflags="-s -w"

build-common:
	node -v && npm -v
	cd frontend && pnpm install && pnpm run build
	go version && go env

build-linux: build-common
	$(GO_BUILD) -o data-viewer backend/cmd/main.go

build-windows: build-common
	$(GO_BUILD) -o data-viewer.exe backend/cmd/main.go

build-mac: build-common
	$(GO_BUILD) -o data-viewer backend/cmd/main.go 