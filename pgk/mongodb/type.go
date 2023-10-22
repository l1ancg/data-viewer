package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Config struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type Client struct {
	client *mongo.Client
	config Config
}
