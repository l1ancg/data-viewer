package application

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
)

type Resource struct {
	Id   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (Resource) TableName() string {
	return "resource"
}

type ResourceService struct {
	pkg.AbstractManager
}

func NewResourceService(db *repository.DB) *ResourceService {
	t := Resource{}
	to := utils.CreateObject("resource", &t)
	dm := ResourceService{AbstractManager: pkg.AbstractManager{
		Name: "resource",
		DB:   db,
		QueryAction: graphql.Fields{
			"resource": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"resources": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Select),
			},
		},
		MutationAction: graphql.Fields{
			"resource": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id", "name", "type", "data"),
				Resolve: utils.CreateSaveResolve(t, db.Save),
			},
		},
		Type: t,
	}}
	return &dm
}
