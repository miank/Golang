package controller

import (
	"example.com/jwt-demo/model"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var reqUser model.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check user exists in DB
	var dbUser model.User
	model.DB.Where("email = ?", reqUser.Email).First(&dbUser)
	if dbUser.Email != "" {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}
	return

}
