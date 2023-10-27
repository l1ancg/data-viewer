package api

import (
	"github.com/google/wire"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/config"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"net/http"
)

func ProvideServer() {
	wire.Build(db.ProvideDB, config.ProvideConfig, log.ProvideLog())
	return {}
	//event := InitializeEvent("hello_world")
	//client := db.New()
	//client.CreateTable(&pkg.Resource{})
	//sc := pkg.InitResourceSchema(client)
	//http.Handle("/graphql", handler.New(&handler.Config{
	//	Schema:   &sc,
	//	Pretty:   true,
	//	GraphiQL: true,
	//}))
	//http.ListenAndServe(":8080", nil)
}
