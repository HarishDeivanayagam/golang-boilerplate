package infrastructure

import (
	"boilerplate/entities"
	"boilerplate/utils/alerts"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) FindOne(email string) (entities.User, error) {
	user := entities.User{}
	emptyUser := entities.User{}
	err := u.db.Get(&user, `SELECT id, name, email, password FROM users WHERE email=?`, email)
	if err != nil {
		return emptyUser, errors.New(alerts.FetchAlert)
	}
	return user, nil
}

func (u *userRepositoryImpl) InsertOne(email string, name string, password string) (entities.User, error) {
	emptyUser := entities.User{}
	res, err := u.db.Exec(`INSERT INTO users (name,email,password) VALUES (?,?,?)`, name, email, password)
	if err != nil {
		return emptyUser, errors.New(alerts.InsertAlert)
	}
	insertId, err := res.LastInsertId()
	if err != nil {
		return emptyUser, errors.New(alerts.InsertAlert)
	}
	user := entities.User{ID: insertId, Name: name, Email: email, Password: password, IsEmailVerified: false, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	return user, nil
}
