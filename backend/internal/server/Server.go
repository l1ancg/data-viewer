package server

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/backend/internal/service"
	"github.com/l1ancg/data-viewer/backend/pkg/config"
	"net/http"
)

type Server struct {
	Config  *config.Config
	Service *service.Service
}

func ServerProvider(config *config.Config, service *service.Service) *Server {
	return &Server{Config: config, Service: service}
}

func (server *Server) Run() {
	sc := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "QUERY",
			Fields: server.Service.QueryActions(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "MUTATION",
			Fields: server.Service.MutationActions(),
		}),
	}
	s, _ := graphql.NewSchema(sc)
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &s,
		Pretty:   true,
		GraphiQL: true,
	}))
	http.ListenAndServe(fmt.Sprintf(":%s", server.Config.Server.Port), nil)
}
