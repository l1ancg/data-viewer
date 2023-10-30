package service

import (
	"github.com/graphql-go/graphql"
	"github.com/l1ancg/data-viewer/backend/pkg"
	"github.com/l1ancg/data-viewer/backend/pkg/db"
	"github.com/l1ancg/data-viewer/backend/pkg/utils"
	"gorm.io/gorm"
)

type Column struct {
	gorm.Model
	Id         int    `json:"id"`
	ResourceId int    `json:"resourceId"`
	Name       string `json:"name"`
	Label      string `json:"label"`
	DataType   string `json:"dataType"`
	OrderBy    string `json:"orderBy"`
	Display    bool   `json:"display"`
	Condition  bool   `json:"condition"`
}

type ColumnService struct {
	pkg.AbstractManager
}

func NewColumnService(db *db.DB) *ColumnService {
	//createTable(db)
	t := Column{}
	to := utils.CreateObject(t)

	dm := ColumnService{AbstractManager: pkg.AbstractManager{
		Name: "column",
		DB:   db,
		QueryAction: map[string]*graphql.Field{
			"column": {
				Type:    to,
				Args:    utils.CreateArguments(t, "id"), // todo select by resource id
				Resolve: utils.CreateGetResolve(t, db.First),
			},
			"columns": {
				Type:    graphql.NewList(to),
				Resolve: utils.CreateListResolve(t, db.Find), // todo order by
			},
		},
		MutationAction: map[string]*graphql.Field{
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

//func createTable2(db *sql.DB) {
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
