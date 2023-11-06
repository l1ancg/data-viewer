package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/internal/repository"
	"github.com/l1ancg/data-viewer/pkg"
	"github.com/l1ancg/data-viewer/pkg/utils"
)

type Column struct {
	Id        int    `json:"id"  gorm:"primarykey"`
	ViewId    int    `json:"viewId"`
	DictId    int    `json:"dictId"`
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	OrderBy   string `json:"orderBy"`
	Display   bool   `json:"display"`
	Condition bool   `json:"condition"`
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
				Resolve: utils.CreateListResolve(t, db.Select),
			},
		},
		MutationAction: graphql.Fields{
			"column": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id", "name", "label", "dataType", "orderBy", "display", "condition"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
		},
		Type: t,
	}}
	return &dm
}
