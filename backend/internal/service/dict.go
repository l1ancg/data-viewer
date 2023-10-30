package service

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
	"gorm.io/gorm"
)

type Dict struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DictDetail struct {
	gorm.Model
	ID      int    `json:"id"`
	GroupId int    `json:"groupId"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

type DictService struct {
	pkg.AbstractManager
}

type DictDetailService struct {
	pkg.AbstractManager
}

func NewDictService(db *db.DB) *DictService {
	//createTable(db)
	t := Dict{}
	to := utils.CreateObject(t)

	// todo query sub table
	// todo saev sub table
	dm := DictService{AbstractManager: pkg.AbstractManager{
		Name: "dict",
		DB:   db,
		QueryAction: map[string]*graphql.Field{
			"dict": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"), // todo select by resource id
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"dicts": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Find), // todo order by
			},
		},
		MutationAction: map[string]*graphql.Field{
			"dict": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateSaveResolve(t, db.Save), // todo save dictDetail
			},
		},
		Type: t,
	}}
	return &dm
}

func NewDictDetailService(db *db.DB) *DictDetailService {
	//createTable(db)
	t := Dict{}
	to := utils.CreateObject(t)

	// todo query sub table
	// todo saev sub table
	dm := DictDetailService{AbstractManager: pkg.AbstractManager{
		Name: "dictDetail",
		DB:   db,
		QueryAction: map[string]*graphql.Field{
			"dictDetails": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Find), // todo order by & query by groupId
			},
		},
		MutationAction: map[string]*graphql.Field{
			"dictDetail": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
		},
		Type: t,
	}}
	return &dm
}
