package user

import (
	"fmt"
	"lab1/db"
	le "lab1/local-errors"
	"lab1/models/category"
	"lab1/models/record"
	"os/user"
)

const StructName = "user"

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Records    []record.Record     `gorm:"foreignKey:UserID"`
	Categories []category.Category `gorm:"foreignKey:CreatedBy"`
}

func Create(name string) (User, error) {
	connect, err := db.Connect()
	if err != nil {
		return User{}, fmt.Errorf("create user on connect: %w", err)
	}

	u := User{
		Name: name,
	}

	tx := connect.Create(&u)
	if tx.Error != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	return u, nil
}

func GetAll() ([]User, error) {
	var result []User

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get users on connect: %w", err)
	}

	tx := connect.Model(user.User{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return result, nil
}

func GetByID(id int) (User, error) {
	connect, err := db.Connect()
	if err != nil {
		return User{}, fmt.Errorf("get user by id on connect: %w", err)
	}

	var u User
	tx := connect.First(&u, id)
	if tx.Error != nil {
		return User{}, fmt.Errorf("get user by id: %w", err)
	}

	return User{}, le.NotFound(StructName)
}
