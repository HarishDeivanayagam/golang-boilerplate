package controllers

import (
	"boilerplate/dtos"
	"boilerplate/services"
	"boilerplate/utils/alerts"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type usersController struct {
	userService services.UserService
}

func NewUsersController(server *gin.Engine, userService services.UserService) {
	u := &usersController{ userService: userService}
	server.POST("/api/v1/users/login", u.login)
	server.POST("/api/v1/users/signup", u.signUp)
}

func (u *usersController) login(c *gin.Context) {
	var postUser dtos.LoginUserDTO
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, alerts.BadRequest)
		return
	}
	err = json.Unmarshal(data, &postUser)
	if err != nil {
		c.String(http.StatusBadRequest, alerts.BadRequest)
		return
	}
	res, err := u.userService.Login(postUser.Email, postUser.Password)
	if err != nil {
		c.String(http.StatusBadRequest, alerts.BadRequest)
		return
	}
	resp := dtos.JWTKey{Email: postUser.Email, Key: res}
	c.JSON(http.StatusOK, resp)
}

func (u *usersController) signUp(c *gin.Context) {
	var postUser dtos.CreateUserDTO
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, alerts.BadRequest)
		return
	}
	err = json.Unmarshal(data, &postUser)
	if err != nil {
		c.String(http.StatusBadRequest, alerts.BadRequest)
		return
	}
	res, err := u.userService.SignUp(postUser.Name, postUser.Email, postUser.Password)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
