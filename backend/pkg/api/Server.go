package api

import (
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
	"log"
	"net/http"
)

func ProvideServer() {
	client := db.New()
	client.CreateTable(&pkg.Resource{})
	sc := pkg.InitResourceSchema(client)
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &sc,
		Pretty:   true,
		GraphiQL: true,
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
