package mysql

import (
	"database/sql"
	"sync"
)

type Config struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name"`
}

type Client struct {
	db     *sql.DB
	config Config
	mutex  sync.Mutex
}
