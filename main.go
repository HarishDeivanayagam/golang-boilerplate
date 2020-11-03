package main

import (
	"clean-architecture/controllers"
	"clean-architecture/di/container"

	"github.com/gin-gonic/gin"
)

func main() {
	container.UseContainer()
	server := gin.New()
	controllers.UseRoutes(server)
	server.Run(":8080")
}
