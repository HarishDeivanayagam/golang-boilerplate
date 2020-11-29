package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type demoController struct {}

func NewDemoController(server *gin.Engine) {
	d := &demoController{}
	server.GET("/api/v1/demo", d.getDemo)
}

func (d *demoController) getDemo(c *gin.Context) {
	c.String(http.StatusOK, "Demo")
	return
}
