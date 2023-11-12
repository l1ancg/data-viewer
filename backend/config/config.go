package config

import (
	"fmt"
	"os"

	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"gopkg.in/ini.v1"
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
}

func NewConfig() *Config {
	return loadConfigFile("data-viewer.ini")
}

func loadConfigFile(name string) *Config {
	cfg, err := ini.Load(name)
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
	log.Logger.Infoln("config load done")
	return config
}
