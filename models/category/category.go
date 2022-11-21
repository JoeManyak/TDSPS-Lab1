package category

import (
	"fmt"
	"lab1/db"
	"lab1/models/structs"
)

func Create(name string, id int) (structs.Category, error) {
	connect, err := db.Connect()
	if err != nil {
		return structs.Category{}, fmt.Errorf("create category on connect: %w", err)
	}

	c := structs.Category{
		Name:      name,
		CreatedBy: id,
	}

	tx := connect.Create(&c)
	if tx.Error != nil {
		return structs.Category{}, fmt.Errorf("create category: %w", tx.Error)
	}

	return c, nil
}

func GetAll() ([]structs.Category, error) {
	var result []structs.Category

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get categories on connect: %w", err)
	}

	tx := connect.Model(structs.Category{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get categories: %w", err)
	}

	return result, nil
}
