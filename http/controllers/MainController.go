package controllers

import (
	"github.com/gin-gonic/gin"
)

type MainController struct {
	Controller
}

func (MainController) Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to talksy roshan",
	})
}

func (MainController) Client(c *gin.Context) {
	c.File("./dist/index.html")
}

func (MainController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ashwat watches fifty shades of grey",
	})
}
