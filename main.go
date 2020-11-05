package main

import (
	"clean-architecture/controllers"
	"clean-architecture/di"
	"clean-architecture/infrastructure"
	"clean-architecture/services"

	"github.com/gin-gonic/gin"
)

func useConfig() {
	di.UseIOC()
	di.AddDependency("BooksRepository", &infrastructure.BooksRepositoryImpl{DB: infrastructure.GetDB()})
	di.AddDependency("BooksService", &services.BooksServiceImpl{})
}

func main() {
	useConfig()
	server := gin.New()
	controllers.UseRoutes(server)
	server.Run(":8080")
}
