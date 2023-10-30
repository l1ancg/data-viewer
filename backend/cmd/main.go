package main

import (
	"github.com/l1ancg/data-viewer/backend/cmd/wire"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	wire.NewServer().Run()
}
