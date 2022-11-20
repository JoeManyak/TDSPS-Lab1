package category

import (
	le "lab1/local-errors"
	"lab1/models/record"
)

var idCount = 0
var categories []Category

const StructName = "category"

type Category struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Records   []record.Record `gorm:"foreignKey:CategoryID"`
	CreatedBy int
}

func init() {
	categories = make([]Category, 0, 10)
}

func Create(name string) {
	categories = append(categories, Category{
		ID:   idCount,
		Name: name,
	})
	idCount++
}

func GetAll() []Category {
	return categories
}

func GetByID(id int) (Category, error) {
	for i := range categories {
		if categories[i].ID == id {
			return categories[i], nil
		}
	}
	return Category{}, le.NotFound(StructName)
}
