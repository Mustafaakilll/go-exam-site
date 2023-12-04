package initialize

import (
	"fmt"

	"github.com/mustafaakilll/exam-site/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=password dbname=examsite port=5432 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	Db.Exec("drop TYPE if exists QuestionType cascade;")
	Db.Exec("CREATE TYPE QuestionType AS ENUM('Multiple Choice','True/False','Open-Ended');")

	Db.Exec("drop TYPE if exists UserType cascade;")
	Db.Exec("CREATE TYPE UserType AS ENUM('admin','teacher','student');")

	if err := Db.AutoMigrate(&models.User{}, &models.Answer{}, &models.Quiz{}, &models.Question{}); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Connection Opened to Database")

}
