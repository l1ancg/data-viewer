package application

import (
	"github.com/l1ancg/data-viewer/backend/internal/repository"
	"github.com/l1ancg/data-viewer/backend/pkg"
)

type Service struct {
	Services *[]pkg.AbstractManager
}

func (s *Service) Names() []string {
	var r []string
	for _, v := range *s.Services {
		r = append(r, v.Name)
	}
	return r
}

func ServiceProvider(db *repository.DB) *Service {
	am := &[]pkg.AbstractManager{
		NewResourceService(db).AbstractManager,
		NewViewService(db).AbstractManager,
		NewColumnService(db).AbstractManager,
		NewDictService(db).AbstractManager,
		NewDictDetailService(db).AbstractManager,
	}
	return &Service{
		Services: am,
	}
}
