package entities

import "github.com/pborman/uuid"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Telephone string `json:"telephone"`
	isAdmin   bool   `json:"isAdmin"`
}

func NewUser() *User {
	User := User{
		ID: uuid.New(),
	}

	return &User
}
