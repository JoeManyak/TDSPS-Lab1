package category

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lab1/db"
	"lab1/models/record"
)

const StructName = "category"

type Category struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Records   []record.Record `gorm:"foreignKey:CategoryID"`
	CreatedBy int             `gorm:"default:null"`
}

func Create(name string, id int) (Category, error) {
	connect, err := db.Connect()
	if err != nil {
		return Category{}, fmt.Errorf("create category on connect: %w", err)
	}

	c := Category{
		Name:      name,
		CreatedBy: id,
	}

	tx := connect.Create(&c)
	if tx.Error != nil {
		return Category{}, fmt.Errorf("create category: %w", tx.Error)
	}

	return c, nil
}

func GetAll() ([]Category, error) {
	var result []Category

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get categories on connect: %w", err)
	}

	tx := connect.Model(Category{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get categories: %w", err)
	}

	return result, nil
}

var ForbiddenCategory = errors.New("cannot use such category")

func CheckCategory(connect *gorm.DB, userID, categoryID int) bool {
	var c Category

	tx := connect.
		Model(&Category{}).
		Where("id = ? AND (created_by = ? OR created_by = NULL)", categoryID, userID).
		Find(&c)

	if c.ID != categoryID || tx.Error != nil {
		return false
	}

	return true
}
