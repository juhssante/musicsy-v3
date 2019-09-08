package controllers

import (
	"App/database/models"
	"App/utils"

	"github.com/gin-gonic/gin"
)

// UserController is a controller for all user operations
type UserController struct {
	Controller
}

// List retuns a list of all users
func (controller UserController) List(c *gin.Context) {
	users := models.UserRepository{}.FindAll()
	c.JSON(200, users)
}

// Post adds a user
func (controller UserController) Post(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	utils.HttpError(c, 400, err)
	user.Store()
	utils.HttpError(c, 500, err)
	c.JSON(200, user)
}

// Patch updates a user
func (controller UserController) Patch(c *gin.Context) {
	if value, ok := c.Get("user"); ok {
		user := value.(models.User)
		var userPatch models.UserPatch
		err := c.Bind(&userPatch)
		utils.HttpError(c, 400, err)
		user.ApplyPatch(userPatch)
		user.Store()
		c.JSON(202, user)
	}

}

// Get returns a user from id
func (controller UserController) Get(c *gin.Context) {
	if value, ok := c.Get("user"); ok {
		user := value.(models.User)
		c.JSON(200, user)
	}
}

// Delete deletes a user
func (controller UserController) Delete(c *gin.Context) {
	if value, ok := c.Get("user"); ok {
		user := value.(models.User)
		err := user.Drop()
		utils.HttpError(c, 500, err)
		c.JSON(204, user)
	}
}
