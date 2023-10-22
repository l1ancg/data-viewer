package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient(configJson string) (*Client, error) {
	config := Config{
		Host: "localhost",
		Port: 27017,
	}
	err := json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.DBName)

	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = db.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	client := &Client{client: db, config: config}

	return client, nil
}

func (c *Client) Query(ql string) ([]map[string]interface{}, error) {
	return nil, nil
}
