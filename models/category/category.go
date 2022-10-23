package category

import le "lab1/local-errors"

var idCount = 0
var categories []Category

const StructName = "category"

type Category struct {
	ID   int
	Name string
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
