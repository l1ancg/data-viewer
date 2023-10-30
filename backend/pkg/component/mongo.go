package component

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type MongoClient struct {
	client *mongo.Client
	config MongoConfig
}

func NewMongoClient(configJson string) (*MongoClient, error) {
	config := MongoConfig{
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

	client := &MongoClient{client: db, config: config}

	return client, nil
}

func (c *MongoClient) Init(data string) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MongoClient) Destroy(ql string) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MongoClient) Query(ql string) ([]map[string]interface{}, error) {
	return nil, nil
}
