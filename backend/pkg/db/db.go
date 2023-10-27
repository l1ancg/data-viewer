package db

import (
	"github.com/l1ancg/data-viewer/backend/pkg/config"
	"github.com/l1ancg/data-viewer/backend/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func ProvideDB(config *config.Config, logger *log.Logger) *DB {
	logger.Info("init sqlite3: %s", config.Sqlite3.File)
	d, err := gorm.Open(sqlite.Open(config.Sqlite3.File), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{db: d}
}

func (db *DB) CreateTable(dst ...interface{}) *DB {
	for i := range dst {
		if !db.db.Migrator().HasTable(dst[i]) {
			if err := db.db.Migrator().CreateTable(dst); err != nil {
				panic(err)
			}
		}
	}

	return db
}
func (db *DB) Save(value interface{}) *DB {
	db.db.Save(value)
	return db
}
func (db *DB) Select(query interface{}, args ...interface{}) *DB {
	db.db.Select(query, args)
	return db
}
func (db *DB) First(query interface{}, conds ...interface{}) *DB {
	db.db.First(query, conds)
	return db
}
func (db *DB) Find(dest interface{}, conds ...interface{}) *DB {
	db.db.Find(dest, conds)
	return db
}
func (db *DB) Update(column string, value interface{}) *DB {
	db.db.Update(column, value)
	return db
}
func (db *DB) Updates(values interface{}) *DB {
	db.db.Updates(values)
	return db
}
func (db *DB) Delete(value interface{}, conds ...interface{}) *DB {
	db.db.Delete(value, conds)
	return db
}
func (db *DB) Model(value interface{}) *DB {
	db.db.Model(value)
	return db
}
