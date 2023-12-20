//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/internal/application"
	"github.com/l1ancg/data-viewer/backend/internal/handler"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/internal/server"
	"github.com/l1ancg/data-viewer/backend/pkg/connect"
)

func NewServer() *server.Server {
	panic(wire.Build(
		config.NewConfig,
		repository.DatabaseProvider,
		application.ApplicationProvider,
		handler.HandlerSet,
		connect.ConnectProvider,
		server.HttpServerProvider,
	))
	return &server.Server{}
}
