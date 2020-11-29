package services

import (
	"boilerplate/entities"
	"boilerplate/repository"
	"boilerplate/utils/alerts"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserServiceImpl(userRepo repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{userRepo: userRepo}
}

func (u *userServiceImpl) Login(email string, password string) (string, error) {
	user, err := u.userRepo.FindOne(email)
	if err != nil {
		return "", errors.New(alerts.BadRequest)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New(alerts.BadRequest)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenString, err := token.SignedString([]byte("IAMNOTMSDHONI7!"))
	if err != nil {
		return "", errors.New(alerts.BadRequest)
	}
	return tokenString, nil
}

func (u *userServiceImpl) SignUp(name string, email string, password string) (entities.User, error) {
	emptyUser := entities.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	res, err := u.userRepo.InsertOne(email, name, string(hashedPassword))
	if err != nil {
		return emptyUser, errors.New(alerts.CreationAlert)
	}
	return res, nil
}
