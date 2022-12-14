package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jfirme-sys/go-auth-docker.git/database"
	"github.com/jfirme-sys/go-auth-docker.git/models"
	"github.com/jfirme-sys/go-auth-docker.git/util"
)

func CreateUser(c *gin.Context) {
	db := database.GetDB()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = util.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
