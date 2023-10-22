package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type QueryResult struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

func main() {
	router := gin.Default()

	// POST /query
	router.POST("/query", func(c *gin.Context) {
		// Parse request body
		var query string
		if err := c.ShouldBindJSON(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Execute query
		db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database_name")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		var result QueryResult
		result.Columns = columns

		for rows.Next() {
			row := make([]interface{}, len(columns))
			for i := range columns {
				row[i] = new(interface{})
			}
			if err := rows.Scan(row...); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			for i := range columns {
				row[i] = *(row[i].(*interface{}))
			}
			result.Rows = append(result.Rows, row)
		}

		// Return result
		c.JSON(http.StatusOK, result)
	})

	// Run server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
