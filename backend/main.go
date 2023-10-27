package main

import (
	"github.com/l1ancg/data-viewer/backend/pkg/api"
	"github.com/l1ancg/data-viewer/backend/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type GlobalEnv struct {
}

func main() {
	config.ProvideConfig()
	api.ProvideServer()
}
