package service

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type ResourceService struct {
	pkg.AbstractManager
}

func NewResourceService(db *db.DB) *ResourceService {
	//createTable(db)
	t := Resource{}
	to := utils.CreateObject(t)
	dm := ResourceService{AbstractManager: pkg.AbstractManager{
		Name: "resource",
		DB:   db,
		QueryAction: map[string]*graphql.Field{
			"resource": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"),
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"resources": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Find), // todo order by
			},
		},
		MutationAction: map[string]*graphql.Field{
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

//func createTable(db *gorm.DB) {
//	createStudentTableSQL := `
//	CREATE TABLE "resource" (
//	  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
//	  "name" TEXT,
//	  "type" TEXT,
//	  "data" TEXT
//	);
//`
//
//	log.Println("Create resource table...")
//	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	statement.Exec() // Execute SQL Statements
//	log.Println("resource table created")
//}
