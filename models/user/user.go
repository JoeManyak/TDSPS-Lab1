package user

import le "lab1/local-errors"

var idCount = 0
var users []User

const StructName = "user"

type User struct {
	ID   int
	Name string
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
