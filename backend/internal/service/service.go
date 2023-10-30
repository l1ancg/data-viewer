package service

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
)

type Service struct {
	Services *[]pkg.AbstractManager
}

func (s *Service) QueryActions() (r map[string]*graphql.Field) {
	ss := s.Services
	for _, v := range *ss {
		for k, v := range v.QueryAction {
			r[k] = v
		}
	}
	return
}

func (s *Service) MutationActions() (r map[string]*graphql.Field) {
	ss := s.Services
	for _, v := range *ss {
		for k, v := range v.MutationAction {
			r[k] = v
		}
	}
	return
}

func ServiceProvider(db *db.DB) *Service {
	am := &[]pkg.AbstractManager{
		NewResourceService(db).AbstractManager,
		NewColumnService(db).AbstractManager,
		NewDictService(db).AbstractManager,
		NewDictDetailService(db).AbstractManager,
	}
	return &Service{
		Services: am,
	}
}
