package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initialize "github.com/mustafaakilll/exam-site/initialize"
	"github.com/mustafaakilll/exam-site/models"
)

func GetLessons(c *gin.Context) {
	var lessons []models.Lesson
	initialize.Db.Find(&lessons)
	c.JSON(http.StatusOK, gin.H{"data": lessons})
}

func CreateLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initialize.Db.Create(&lesson)
	c.JSON(http.StatusOK, gin.H{"data": lesson})
}
