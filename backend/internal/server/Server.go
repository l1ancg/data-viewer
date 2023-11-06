package server

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/config"
	"github.com/l1ancg/data-viewer/internal/application"
	"github.com/l1ancg/data-viewer/pkg"
	"github.com/l1ancg/data-viewer/pkg/log"
)

type Server struct {
	Config *config.Config
}

type GraphQLHandler struct {
	pkg.AbstractHandler
}

func GraphQLHandlerProvider(service *application.Service) *GraphQLHandler {
	queryFields := graphql.Fields{}
	mutationFields := graphql.Fields{}
	for _, manager := range *service.Services {
		for key, val := range manager.QueryAction {
			queryFields[key] = val
		}
		for key, val := range manager.MutationAction {
			mutationFields[key] = val
		}
	}
	sc := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queryFields,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutationFields,
		}),
	}
	schema, _ := graphql.NewSchema(sc)
	handler := GraphQLHandler{AbstractHandler: pkg.AbstractHandler{
		Path: "/graphql",
		Handler: handler.New(&handler.Config{
			Schema:     &schema,
			Pretty:     true,
			Playground: true,
		}),
	}}
	log.Logger.Infoln("graphql query handler init success:", service.Names())
	return &handler
}

func HttpServerProvider(config *config.Config, handler *GraphQLHandler) *Server {
	http.Handle(handler.Path, handler.Handler)
	return &Server{Config: config}
}

func (server *Server) Run() {
	log.Logger.Infoln("server run on port:", server.Config.Server.Port)
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", server.Config.Server.Port), nil)
}
