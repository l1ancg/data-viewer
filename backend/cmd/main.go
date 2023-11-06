package main

import (
	"github.com/l1ancg/data-viewer/cmd/wire"
	"github.com/l1ancg/data-viewer/pkg/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Logger.Infoln("start server")
	wire.NewServer().Run()
	log.Logger.Infoln("start server")
}
