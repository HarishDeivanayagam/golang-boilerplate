package app

import (
	"boilerplate/controllers"
	"boilerplate/infrastructure"
	"boilerplate/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type app struct {}

func New() *app {
	return &app{}
}

func (a *app) useConfig() {
	server := gin.New()
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/boilerplate")
	if err != nil {
		log.Println("Database Initialization Failed")
		return
	}
	userRepo := infrastructure.NewUserRepositoryImpl(db)
	userService := services.NewUserServiceImpl(userRepo)
	controllers.NewDemoController(server)
	controllers.NewUsersController(server, userService)
	log.Println("App listening on http://localhost:8080/api/v1/")
	server.Run(":8080")
}

func (a *app) Init()  {
	a.useConfig()
}
