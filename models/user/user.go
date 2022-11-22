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

func Equal(u1, u2 structs.User) bool {
	return u1.ID == u2.ID && u1.Name == u2.Name
}

func GetByID(id int) (structs.User, error) {
	var result structs.User

	connect, err := db.Connect()
	if err != nil {
		return structs.User{}, fmt.Errorf("get user by id on connect: %w", err)
	}

	tx := connect.Model(result).Where(structs.User{ID: id}).Find(&result)
	if tx.Error != nil {
		return result, fmt.Errorf("get user by id: %w", err)
	}

	return result, nil
}
