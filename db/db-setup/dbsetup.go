package db_setup

import (
	"fmt"
	"gorm.io/gorm"
	"lab1/db"
	"lab1/models/category"
	"lab1/models/record"
	"lab1/models/user"
	"os"
)

func Clear() error {
	return os.Remove(db.Filename)
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&category.Category{},
		&user.User{},
		&record.Record{},
	)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}

func CreateDefaultData(db *gorm.DB) error {
	for _, v := range getDefaultUsers() {
		tx := db.Create(&v)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func getDefaultUsers() []user.User {
	return []user.User{
		{
			Name: "Joe",
		},
		{
			Name: "Freddy",
		},
		{
			Name: "Leo",
		},
	}
}
