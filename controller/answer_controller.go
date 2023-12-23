package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initialize "github.com/mustafaakilll/exam-site/initialize"
	"github.com/mustafaakilll/exam-site/models"
)

func GetAnswers(c *gin.Context) {
	var answers []models.Answer
	initialize.Db.Find(&answers)
	c.JSON(http.StatusOK, gin.H{"data": answers})
}

func CreateAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initialize.Db.Create(&answer)
	c.JSON(http.StatusOK, gin.H{"data": answer})
}
