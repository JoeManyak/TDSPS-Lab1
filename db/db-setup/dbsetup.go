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
	var lines int64
	if db.Model(&user.User{}).Count(&lines); lines == 0 {
		for _, v := range getDefaultUsers() {
			tx := db.Create(&v)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	if db.Model(&category.Category{}).Count(&lines); lines == 0 {
		for _, v := range getDefaultCategories() {
			tx := db.Create(&v)
			if tx.Error != nil {
				return tx.Error
			}
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

func getDefaultCategories() []category.Category {
	return []category.Category{
		{
			Name: "Default",
		},
		{
			Name: "Premium",
		},
	}
}
