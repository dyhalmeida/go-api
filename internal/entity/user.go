package entity

import (
	"github.com/dyhalmeida/go-apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id       entity.ID
	name     string
	email    string
	password string
}

func NewUser(name, email, password string) (*user, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &user{
		id:       entity.NewID(),
		name:     name,
		email:    email,
		password: string(hash),
	}, nil
}

func (u *user) GetID() entity.ID {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}
