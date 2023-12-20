package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
)

type View struct {
	Id         int    `json:"id"  gorm:"primarykey"`
	Name       string `json:"name"`
	ResourceId int    `json:"resourceId"`
	Ql         string `json:"ql"`
	Options    string `json:"options"`
}

func (View) TableName() string {
	return "view"
}

type ViewService struct {
	pkg.AbstractManager
}

func NewViewService(db *repository.Database) *ViewService {
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
				Args:    utils.CreateArguments(t, "id", "name", "resourceId", "ql", "options"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
			"delete": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateDeleteResolve(t, db.Delete),
			},
		},
		Type: t,
	}}
	return &dm
}
