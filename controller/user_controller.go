package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initialize "github.com/mustafaakilll/exam-site/initialize"
	"github.com/mustafaakilll/exam-site/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	initialize.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initialize.Db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}