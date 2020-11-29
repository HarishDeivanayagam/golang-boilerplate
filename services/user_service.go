package services

import "boilerplate/entities"

type UserService interface {
	Login(email string, password string) (string, error)
	SignUp(name string, email string, password string) (entities.User, error)
}
