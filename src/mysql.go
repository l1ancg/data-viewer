package src

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type MySQLClient struct {
	db     *sql.DB
	config MySQLConfig
	mutex  sync.Mutex
}

func (c *MySQLClient) Init(configData string) (*MySQLClient, error) {
	config := MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
	}
	err := json.Unmarshal([]byte(configData), &config)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, err
	}

	client := &MySQLClient{
		db:     db,
		config: config,
	}

	return client, nil
}

func (c *MySQLClient) Destroy(ql string) {
	//TODO implement me
	panic("implement me")
}

func (c *MySQLClient) Query(query string) ([]map[string]interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	return result, nil
}
