package db_setup

import (
	"fmt"
	"gorm.io/gorm"
	"lab1/db"
	"lab1/models/structs"
	"os"
)

func Clear() error {
	return os.Remove(db.Filename)
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&structs.Category{},
		&structs.User{},
		&structs.Record{},
	)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}

func CreateDefaultData(db *gorm.DB) error {
	var lines int64
	if db.Model(&structs.User{}).Count(&lines); lines == 0 {
		for _, v := range getDefaultUsers() {
			tx := db.Create(&v)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	if db.Model(&structs.Category{}).Count(&lines); lines == 0 {
		for _, v := range getDefaultCategories() {
			tx := db.Create(&v)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}
	return nil
}

func getDefaultUsers() []structs.User {
	return []structs.User{
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

func getDefaultCategories() []structs.Category {
	return []structs.Category{
		{
			Name: "Default",
		},
		{
			Name: "Premium",
		},
	}
}
