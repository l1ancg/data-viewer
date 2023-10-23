package src

import (
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
	"log"
)

type Resource struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data string `json:"data"`
}

var resource = graphql.NewObject(graphql.ObjectConfig{
	Name: "resource",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func createMutation(db *sql.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "resource",
		Fields: graphql.Fields{
			"saveResource": &graphql.Field{
				Type:        resource,
				Description: "Update or save new resource",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"data": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					_type, _ := params.Args["type"].(string)
					data, _ := params.Args["data"].(string)

					insertStudentSQL := `INSERT INTO resource(id, name, type, data) VALUES (?, ?, ?, ?)`
					statement, err := db.Prepare(insertStudentSQL)
					if err != nil {
						return nil, err
					}
					_, err = statement.Exec(id, name, _type, data)
					if err != nil {
						return nil, err
					}
					return &Resource{
						Id:   id,
						Name: name,
						Type: _type,
						Data: data,
					}, nil
				},
			},
		},
	})
}

func createQuery(db *sql.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "resource",
		Fields: graphql.Fields{
			"resource": &graphql.Field{
				Type:        resource,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, isOK := params.Args["id"].(string)
					if isOK {
						rows, err := db.Query("SELECT id, name, type, data FROM resource WHERE id = ?", id)
						if err != nil {
							return nil, err
						}
						if !rows.Next() {
							return nil, errors.New("没找到")
						}
						var r Resource
						err = rows.Scan(&r.Id, &r.Name, &r.Type, &r.Data)
						if err != nil {
							return nil, err
						}
						return r, nil
					}
					return nil, errors.New("请指定ID参数")
				},
			},
			"resources": &graphql.Field{
				Type:        graphql.NewList(resource),
				Description: "List of resources",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT ID, NAME, TYPE, DATA FROM main.resource")
					if err != nil {
						return nil, err
					}
					var rs []Resource
					for rows.Next() {
						var r Resource
						err = rows.Scan(&r.Id, &r.Name, &r.Type, &r.Data)
						if err != nil {
							return nil, err
						}
						rs = append(rs, r)
					}
					return rs, nil
				},
			},
		},
	})
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE "resource" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "type" integer,
  "data" TEXT
);`

	log.Println("Create resource table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("resource table created")
}

func InitResourceSchema(db *sql.DB) graphql.Schema {
	//createTable(db)
	sc := graphql.SchemaConfig{
		Query:    createQuery(db),
		Mutation: createMutation(db),
	}
	s, _ := graphql.NewSchema(sc)
	return s
}
