package pkg

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/labstack/echo/v4"
)

type AbstractManager struct {
	Name           string
	DB             *repository.Database
	QueryAction    graphql.Fields
	MutationAction graphql.Fields
	Type           interface{}
}

type AbstractHandler struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

type Client interface {
	Query(typ string, id int, uri string, ql string) ([]map[string]interface{}, error)
}
