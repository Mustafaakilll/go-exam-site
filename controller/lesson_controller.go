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
