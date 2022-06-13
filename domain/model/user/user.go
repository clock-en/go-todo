package user

import (
	"github.com/clock-en/go-todo/domain/model"
)

// User represents entity
type User struct {
	id       model.Id
	name     Name
	email    model.Email
	password model.Password
}
type Users []User

func NewUser(id model.Id, name Name, email model.Email, password model.Password) *User {
	return &User{id: id, name: name, email: email, password: password}
}
