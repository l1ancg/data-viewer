package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type GlobalEnv struct {
	Server  int     `ini:"server"`
	Sqlite3 sqlite3 `ini:"sqlite3"`
}

type server struct {
	Port int `ini:"port"`
}

type sqlite3 struct {
	File string `ini:"file"`
	Name string `ini:"name"`
}

var Env *GlobalEnv

func LoadConfig() {
	cfg, err := ini.Load("data-viewer.ini")
	if err != nil {
		fmt.Printf("Fail to read config file: %v", err)
		os.Exit(1)
	}
	Env = &GlobalEnv{}
	err = cfg.MapTo(Env)
	if err != nil {
		fmt.Printf("Fail to read config file: %v", err)
		os.Exit(1)
	}
	fmt.Println("env config load done: ", Env)
}
