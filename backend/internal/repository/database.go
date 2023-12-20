package repository

import (
	"github.com/l1ancg/data-viewer/backend/config"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func DatabaseProvider(config *config.Config) *Database {
	log.Logger.Infoln("init sqlite3:", config.Sqlite3.File)
	d, err := gorm.Open(sqlite.Open(config.Sqlite3.File), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for _, sql := range sqls {
		d.Exec(sql)
	}

	log.Logger.Infoln("sqlite3 table init done")
	return &Database{db: d}
}

func (database *Database) Save(value interface{}) {
	log.Logger.Infof("save data: %+v", value)
	database.db.Save(value)
}

func (database *Database) Select(dest interface{}, conds ...interface{}) {
	database.db.Order("id").Find(dest, conds...)
	log.Logger.Infof("select result: %v", dest)
}

func (database *Database) First(query interface{}, id int) {
	database.db.First(query, id)
	log.Logger.Infof("first result: %+v", query)
}

func (database *Database) Delete(query interface{}, id int) {
	database.db.Delete(query, id)
	log.Logger.Infof("delete: %+v", id)
}
