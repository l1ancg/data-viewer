package server

import (
	"fmt"
	"github.com/brpaz/echozap"
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/internal/handler"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Config *config.Config
	Engine *echo.Echo
}

func HttpServerProvider(config *config.Config, graphql *handler.GraphQLHandler, query *handler.QueryHandler) *Server {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(echozap.ZapLogger(log.Logger.Desugar()))
	e.Use(middleware.CORS())

	e.Add(query.Method, graphql.Path, graphql.Handler)
	e.Add(query.Method, query.Path, query.Handler)

	e.Static("/", "frontend/dist")

	return &Server{Config: config, Engine: e}
}

func (server *Server) Run() {
	server.Engine.Logger.Fatal(server.Engine.Start(fmt.Sprintf("127.0.0.1:%s", server.Config.Server.Port)))
	log.Logger.Infoln("server run on port:", server.Config.Server.Port)
}
