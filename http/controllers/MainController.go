package controllers

import (
	"github.com/gin-gonic/gin"
)

// MainController is the Main handler for the client and pinger
type MainController struct {
	Controller
}

// Homepage is just a page to make fun of fools
func (MainController) Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to talksy roshan",
	})
}

// Client hydrates and serves the client file
func (MainController) Client(c *gin.Context) {
	c.File("./dist/index.html")
}

// Ping is used for uptime
func (MainController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ashwat watches fifty shades of grey",
	})
}
