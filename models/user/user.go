package user

import (
	"fmt"
	"lab1/db"
	"lab1/models/structs"
)

func Create(name string) (structs.User, error) {
	connect, err := db.Connect()
	if err != nil {
		return structs.User{}, fmt.Errorf("create user on connect: %w", err)
	}

	u := structs.User{
		Name: name,
	}

	tx := connect.Create(&u)
	if tx.Error != nil {
		return structs.User{}, fmt.Errorf("create user: %w", tx.Error)
	}

	return u, nil
}

func GetAll() ([]structs.User, error) {
	var result []structs.User

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get users on connect: %w", err)
	}

	tx := connect.Model(structs.User{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return result, nil
}
