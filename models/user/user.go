package user

import (
	le "lab1/local-errors"
	"lab1/models/category"
	"lab1/models/record"
)

var idCount = 0
var users []User

const StructName = "user"

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Records    []record.Record     `gorm:"foreignKey:UserID"`
	Categories []category.Category `gorm:"foreignKey:CreatedBy"`
}

func init() {
	users = make([]User, 0, 10)
}

func Create(name string) {
	users = append(users, User{
		ID:   idCount,
		Name: name,
	})
	idCount++
}

func GetAll() []User {
	return users
}

func GetByID(id int) (User, error) {
	for i := range users {
		if users[i].ID == id {
			return users[i], nil
		}
	}
	return User{}, le.NotFound(StructName)
}
