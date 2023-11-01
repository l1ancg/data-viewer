package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
)

type View struct {
	Id         int    `json:"id"  gorm:"primarykey"`
	ResourceId int    `json:"resourceId"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
}

func (View) TableName() string {
	return "view"
}

type ViewService struct {
	pkg.AbstractManager
}

func NewViewService(db *repository.DB) *ViewService {
	t := View{}
	to := utils.CreateObject("view", &t)
	dm := ViewService{AbstractManager: pkg.AbstractManager{
		Name: "view",
		DB:   db,
		QueryAction: graphql.Fields{
			"view": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"views": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Select),
			},
		},
		MutationAction: graphql.Fields{
			"view": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id", "resourceId", "name", "desc"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
		},
		Type: t,
	}}
	return &dm
}
