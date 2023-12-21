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
	r.GET("/lessons", controller.GetLessons)

	http.ListenAndServe(":3000", r)
}
