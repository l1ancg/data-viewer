package repository

import (
	"github.com/l1ancg/data-viewer/config"
	"github.com/l1ancg/data-viewer/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func DBProvider(config *config.Config) *DB {
	log.Logger.Infoln("init sqlite3:", config.Sqlite3.File)
	d, err := gorm.Open(sqlite.Open(config.Sqlite3.File), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	d.Exec(initSql)
	log.Logger.Infoln("sqlite3 table init done")
	return &DB{db: d}
}

func (db *DB) Save(value interface{}) {
	log.Logger.Infof("save data: %+v", value)
	db.db.Save(value)
}
func (db *DB) Select(dest interface{}) {
	db.db.Order("id").Find(dest)
	log.Logger.Infof("select result: %v", dest)
}

func (db *DB) First(query interface{}, id int) {
	db.db.First(query, id)
	log.Logger.Infof("first result: %+v", query)
}

var initSql = `
CREATE TABLE IF NOT EXISTS "resource" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT,
    "type" INTEGER,
    "data" TEXT
);


CREATE TABLE IF NOT EXISTS "view" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "resource_id" TEXT,
    "name" TEXT,
    "desc" TEXT
);


CREATE TABLE IF NOT EXISTS "column" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "view_id" TEXT,
    "dict_id" TEXT,
    "name" TEXT,
    "dataType" TEXT,
    "orderBy" TEXT,
    "display" INTEGER,
    "condition" INTEGER
);


CREATE TABLE IF NOT EXISTS "dict" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT
);


CREATE TABLE IF NOT EXISTS "dict_detail" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "dict_id" INTEGER,
    "key" TEXT,
    "value" TEXT
);
`
