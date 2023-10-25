package pkg

import (
	"errors"
	"github.com/graphql-go/graphql"
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

var resourceObject = utils.CreateObject(Resource{})

func createMutation(db *db.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "resource",
		Fields: graphql.Fields{
			"saveResource": &graphql.Field{
				Type:        resourceObject,
				Description: "Update or save new resource",
				Args:        utils.CreateArguments(Resource{}, "id", "name", "type", "data"),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					_type, _ := params.Args["type"].(string)
					data, _ := params.Args["data"].(string)
					r := &Resource{
						Id:   id,
						Name: name,
						Type: _type,
						Data: data,
					}
					db.Save(r)
					return r, nil
				},
			},
		},
	})
}

func createQuery(db *db.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "resource",
		Fields: graphql.Fields{
			"resource": &graphql.Field{
				Type:        resourceObject,
				Description: "",
				Args:        utils.CreateArguments(Resource{}, "id"),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, isOK := params.Args["id"].(string)
					if isOK {
						var r Resource
						db.First(&r, id)
						return r, nil
					}
					return nil, errors.New("请指定ID参数")
				},
			},
			"resources": &graphql.Field{
				Type:        graphql.NewList(resourceObject),
				Description: "List of resources",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var rs []Resource
					db.Find(rs)
					return rs, nil
				},
			},
		},
	})
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

func InitResourceSchema(db *db.DB) graphql.Schema {
	//createTable(db)
	sc := graphql.SchemaConfig{
		Query:    createQuery(db),
		Mutation: createMutation(db),
	}
	s, _ := graphql.NewSchema(sc)
	return s
}
