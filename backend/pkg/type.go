package pkg

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/l1ancg/data-viewer/internal/repository"
	"github.com/l1ancg/data-viewer/pkg/component"
)

//type Action map[string]*graphql.Field

type AbstractManager struct {
	Name           string
	DB             *repository.DB
	QueryAction    graphql.Fields
	MutationAction graphql.Fields
	Type           interface{}
}

type AbstractHandler struct {
	Path    string
	Handler *handler.Handler
	// auth...
}

//func (a *Action) Names() []string {
//	var r []string
//	for k := range *a {
//		r = append(r, k)
//	}
//	return r
//}

type Connect interface {
	Init(data string) (*component.MySQLClient, error)
	Destroy(ql string)
	Query(ql string) ([]map[string]interface{}, error)
}
