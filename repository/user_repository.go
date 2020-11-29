package repository

import (
	"boilerplate/entities"
)

type UserRepository interface {
	FindOne(email string) (entities.User, error)
	InsertOne(email string, name string, password string) (entities.User, error)
}
