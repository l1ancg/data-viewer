package pkg

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/pkg/component"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
)

type AbstractManager struct {
	Name           string
	DB             *db.DB
	QueryAction    map[string]*graphql.Field
	MutationAction map[string]*graphql.Field
	Type           interface{}
}

type Connect interface {
	Init(data string) (*component.MySQLClient, error)
	Destroy(ql string)
	Query(ql string) ([]map[string]interface{}, error)
}
