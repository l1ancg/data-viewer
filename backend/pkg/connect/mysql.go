package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	dbs map[int]*gorm.DB
}

func NewMySQLClient() *MySQLClient {
	return &MySQLClient{}
}

func (client *MySQLClient) Query(_ string, id int, uri string, sql string) ([]map[string]interface{}, error) {
	if _, ok := client.dbs[id]; !ok {
		db, err := gorm.Open(mysql.Open(uri), &gorm.Config{QueryFields: true})
		if err != nil {
			return nil, err
		}
		if client.dbs == nil {
			client.dbs = make(map[int]*gorm.DB)
		}
		client.dbs[id] = db
	}
	var result []map[string]interface{}
	client.dbs[id].Raw(sql).Scan(&result)
	return result, nil
}
