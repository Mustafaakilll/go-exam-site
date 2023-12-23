package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafaakilll/exam-site/controller"
	initialize "github.com/mustafaakilll/exam-site/initialize"
)

func init() {
	initialize.Connect()
}

func main() {

	r := gin.New()
	r.GET("/answers", controller.GetAnswer)
	r.POST("/answer", controller.CreateAnswer)

	r.GET("/lessons", controller.GetLesson)
	r.POST("/lesson", controller.CreateLesson)

	r.GET("/questions", controller.GetQuestion)
	r.POST("/question", controller.CreateQuestion)

	r.GET("/quizs", controller.GetQuizzes)
	r.POST("/quiz", controller.CreateQuiz)

	r.GET("/users", controller.GetUser)
	r.POST("/user", controller.CreateUser)

	http.ListenAndServe(":3000", r)
}
