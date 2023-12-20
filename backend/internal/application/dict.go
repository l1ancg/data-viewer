package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
)

type Dict struct {
	Id   int    `json:"id"  gorm:"primarykey"`
	Name string `json:"name"`
}

type DictDetail struct {
	Id     int    `json:"id"  gorm:"primarykey"`
	DictId int    `json:"dictId"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func (Dict) TableName() string {
	return "dict"
}

func (DictDetail) TableName() string {
	return "dict_detail"
}

type DictService struct {
	pkg.AbstractManager
}

type DictDetailService struct {
	pkg.AbstractManager
}

func NewDictService(db *repository.Database) *DictService {
	t := Dict{}
	to := utils.CreateObject("dict", &t)

	dm := DictService{AbstractManager: pkg.AbstractManager{
		Name: "dict",
		DB:   db,
		QueryAction: graphql.Fields{
			"dict": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"dicts": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Select),
			},
		},
		MutationAction: graphql.Fields{
			"dict": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
		},
		Type: t,
	}}
	return &dm
}

func NewDictDetailService(db *repository.Database) *DictDetailService {
	t := Dict{}
	to := utils.CreateObject("dictDetail", &t)

	dm := DictDetailService{AbstractManager: pkg.AbstractManager{
		Name: "dictDetail",
		DB:   db,
		QueryAction: graphql.Fields{
			"dictDetails": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Select),
			},
		},
		MutationAction: graphql.Fields{
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
