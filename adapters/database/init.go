package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB gorm db
type DB = gorm.DB

// Database extend gorm.DB
type Database struct {
	*DB
}

// Init init sqlite database
func Init() (db *Database, err error) {
	gormDB, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		err = errors.New("failed to connect database")
	}

	db = &Database{gormDB}

	return db, err
}
