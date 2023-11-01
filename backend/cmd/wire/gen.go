//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/internal/server"
)

func NewServer() *server.Server {
	panic(wire.Build(
		config.NewConfig,
		repository.DBProvider,
		//application.ServiceProvider,
		server.GraphQLHandlerProvider,
		server.HttpServerProvider,
	))
	return &server.Server{}
}
