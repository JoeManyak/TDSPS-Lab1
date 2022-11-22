package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const Filename = "./data.db"

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(Filename), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}
	return db, nil
}
