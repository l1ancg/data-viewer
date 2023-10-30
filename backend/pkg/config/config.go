package config

import (
	"fmt"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	Server  server  `ini:"server"`
	Sqlite3 sqlite3 `ini:"sqlite3"`
}

type server struct {
	Port string `ini:"port"`
}

type sqlite3 struct {
	File string `ini:"file"`
	Name string `ini:"name"`
}

func NewConfig() *Config {
	cfg, err := ini.Load("data-viewer.ini")
	if err != nil {
		fmt.Printf("Fail to read config file: %v", err)
		os.Exit(1)
	}
	config := &Config{}
	err = cfg.MapTo(config)
	if err != nil {
		fmt.Printf("Fail to read config file: %v", err)
		os.Exit(1)
	}
	log.Info("config load done: %s", config)
	return config
}
