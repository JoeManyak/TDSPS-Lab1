package user

import (
	"fmt"
	"lab1/db"
	"lab1/models/category"
	"lab1/models/record"
)

const StructName = "user"

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Records    []record.Record     `json:",omitempty" gorm:"foreignKey:UserID"`
	Categories []category.Category `json:",omitempty" gorm:"foreignKey:CreatedBy"`
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
		return User{}, fmt.Errorf("create user: %w", tx.Error)
	}

	return u, nil
}

func GetAll() ([]User, error) {
	var result []User

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get users on connect: %w", err)
	}

	tx := connect.Model(User{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return result, nil
}
