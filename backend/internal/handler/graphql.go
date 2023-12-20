package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/backend/internal/application"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"github.com/labstack/echo/v4"
)

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
	h2 := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})
	h := GraphQLHandler{AbstractHandler: pkg.AbstractHandler{
		Method: "POST",
		Path:   "/graphql",
		Handler: func(c echo.Context) error {
			req := c.Request()
			res := c.Response().Writer
			h2.ServeHTTP(res, req)
			return nil
		},
	}}
	log.Logger.Infoln("graphql query handler init success:", service.Names())
	return &h
}
