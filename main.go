package main

import (
	"database/sql"
	"github.com/graphql-go/handler"
	"github.com/l1an10/dataserver/src"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data-viewer.db")
	db.Query("select * from resource")
	if err != nil {
		// handle error
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		// handle error
		panic(err)
	}
	return db
}

func main() {
	db := setupDB()
	defer db.Close()
	sc := src.InitResourceSchema(db)
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &sc,
		Pretty:   true,
		GraphiQL: true,
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
