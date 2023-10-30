//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/l1ancg/data-viewer/backend/internal/server"
	"github.com/l1ancg/data-viewer/backend/internal/service"
	"github.com/l1ancg/data-viewer/backend/pkg/config"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
)

func NewServer() *server.Server {
	panic(wire.Build(
		config.NewConfig,
		db.DBprovider,
		service.ServiceProvider,
		server.ServerProvider,
	))
	return &server.Server{}
}
