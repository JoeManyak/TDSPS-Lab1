package user

var idCount = 0
var users []User

type User struct {
	ID   int
	Name string
}
