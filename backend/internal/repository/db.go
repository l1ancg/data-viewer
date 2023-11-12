package repository

import (
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
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
	for _, sql := range sqls {
		d.Exec(sql)
	}

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
