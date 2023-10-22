package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func NewMySQLClient(configJson string) (*Client, error) {
	config := Config{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
	}
	err := json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		return nil, err
	}

	client := &Client{
		db:     db,
		config: config,
	}

	return client, nil
}

func (c *Client) Query(query string) ([]map[string]interface{}, error) {
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
