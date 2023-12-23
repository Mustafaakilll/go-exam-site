package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initialize "github.com/mustafaakilll/exam-site/initialize"
	"github.com/mustafaakilll/exam-site/models"
)

func GetQuizzes(c *gin.Context) {
	var quizzes []models.Quiz
	initialize.Db.Find(&quizzes)
	c.JSON(http.StatusOK, gin.H{"data": quizzes})
}

func CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initialize.Db.Create(&quiz)
	c.JSON(http.StatusOK, gin.H{"data": quiz})
}
