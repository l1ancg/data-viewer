wire:
	$(shell go env GOPATH)/bin/wire backend/cmd/wire/gen.go

backend-run:
	cd backend
	go run backend/cmd/main.go

frontend-run:
	cd frontend && pnpm install && pnpm run dev

GO_BUILD = go build -ldflags="-s -w"

build-all:
	node -v && npm -v
	go version && go env
	cd frontend && npm install && npm run build
	rm -rf ./bin/data-viewer*
	CGO_ENABLED=0 GOOS=linux $(GO_BUILD) -o bin/data-viewer backend/cmd/main.go
	GOARM=7 GOARCH=arm CGO_ENABLED=0 GOOS=linux $(GO_BUILD) -o bin/data-viewer-arm64 backend/cmd/main.go
	CGO_ENABLED=0 GOOS=darwin $(GO_BUILD) -o bin/data-viewer-darwin backend/cmd/main.go
	GOARCH=arm64 CGO_ENABLED=0 GOOS=darwin $(GO_BUILD) -o bin/data-viewer-darwin-arm64 backend/cmd/main.go
	GOOS=windows CGO_ENABLED=0 $(GO_BUILD) -o bin/data-viewer.exe backend/cmd/main.go
	pwd && ls ./bin -alh