package controllers

import "github.com/gin-gonic/gin"

// UseRoutes contains all the routes for all the controllers
func UseRoutes(server *gin.Engine) {
	server.GET("/", getAllBooks)
}
