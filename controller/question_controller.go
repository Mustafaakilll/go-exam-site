package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initialize "github.com/mustafaakilll/exam-site/initialize"
	"github.com/mustafaakilll/exam-site/models"
)

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	initialize.Db.Find(&questions)
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initialize.Db.Create(&question)
	c.JSON(http.StatusOK, gin.H{"data": question})
}