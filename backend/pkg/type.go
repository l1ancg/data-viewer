package pkg

import "github.com/l1ancg/data-viewer/backend/component"

type Crud interface {
	save() error
	findById() error
	findAll() error
	delete() error
	count() error
}

type Connect interface {
	Init(data string) (*component.MySQLClient, error)
	Destroy(ql string)
	Query(ql string) ([]map[string]interface{}, error)
}
