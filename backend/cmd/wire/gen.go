//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/l1ancg/data-viewer/config"
	"github.com/l1ancg/data-viewer/internal/repository"
	"github.com/l1ancg/data-viewer/internal/server"
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
