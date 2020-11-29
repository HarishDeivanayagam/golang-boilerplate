package entities

import "time"

type User struct {
	ID	int64
	Name string
	Email string
	Password string
	IsEmailVerified bool
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
