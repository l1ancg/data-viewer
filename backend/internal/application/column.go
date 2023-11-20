package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
)

type Column struct {
	Id         int    `json:"id"  gorm:"primarykey"`
	ResourceId int    `json:"resourceId"`
	DictId     int    `json:"dictId"`
	Name       string `json:"name"`
	DataType   string `json:"dataType"`
	OrderBy    string `json:"orderBy"`
	Display    bool   `json:"display"`
	Condition  bool   `json:"condition"`
	Desc       string `json:"desc"`
}

func (Column) TableName() string {
	return "column"
}

type ColumnService struct {
	pkg.AbstractManager
}

func NewColumnService(db *repository.DB) *ColumnService {
	t := Column{}
	to := utils.CreateObject("column", &t)
	dm := ColumnService{AbstractManager: pkg.AbstractManager{
		Name: "column",
		DB:   db,
		QueryAction: graphql.Fields{
			"column": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"columns": {
				Type:    graphql.NewList(to),
				Args:    utils.CreateArguments(t, "resourceId"),
				Resolve: utils.CreateParamListResolve(t, db.Select, "resourceId"),
			},
		},
		MutationAction: graphql.Fields{
			"createColumn": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id", "resourceId", "dictId", "name", "dataType", "orderBy", "display", "condition", "desc"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
			"deleteColumn": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateDeleteResolve(t, db.Delete),
			},
		},
		Type: t,
	}}
	return &dm
}
