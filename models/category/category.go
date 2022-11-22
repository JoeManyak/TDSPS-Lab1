package category

import (
	"fmt"
	"lab1/db"
	"lab1/models/structs"
)

func Parse(data map[string]any) (structs.Category, bool) {
	c := structs.Category{
		Name:      "",
		CreatedBy: 0,
	}
	if name, ok := data["name"].(string); ok {
		c.Name = name
	} else {
		return structs.Category{}, false
	}

	if createdBy, ok := data["created_by"].(float64); ok {
		c.CreatedBy = int(createdBy)
	} else {
		return structs.Category{}, false
	}
	return c, true
}

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
