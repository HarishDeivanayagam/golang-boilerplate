package controllers

import (
	"clean-architecture/di/injector"
	"clean-architecture/services"

	"github.com/gin-gonic/gin"
)

func getAllBooks(c *gin.Context) {
	bs, ok := injector.Inject("BooksService").(services.BooksService)
	if !ok {
		c.String(500, "Unable to Contact Server")
		return
	}

	res := bs.GetAllBooks()
	if res == nil {
		c.String(404, "No Result")
		return
	}
	c.JSON(200, res)
	return
}
