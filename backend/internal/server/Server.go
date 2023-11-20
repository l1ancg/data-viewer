package server

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/internal/application"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
)

type Server struct {
	Config     *config.Config
	MuxHandler *http.Handler
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
	mux := http.NewServeMux()
	mux.Handle(handler.Path, handler.Handler)
	muxHandler := middleware(mux)
	return &Server{Config: config, MuxHandler: &muxHandler}
}

func (server *Server) Run() {
	log.Logger.Infoln("server run on port:", server.Config.Server.Port)
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", server.Config.Server.Port), *server.MuxHandler)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置跨域的响应头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 如果是预检请求，直接返回
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
