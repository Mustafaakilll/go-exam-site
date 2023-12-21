package initialize

import (
	"fmt"

	"github.com/mustafaakilll/exam-site/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const createSQL = `
  insert into public.lessons (lesson_name) Values('Isletim Sistemleri');
  insert into public.lessons (lesson_name) Values('Veri Yapilari');

  INSERT INTO users( username, email, name, surname, password, user_type)
	VALUES ('mustafaakilll', 'mustafa@deneme.com', 'mustafa', 'akil', '123456','admin');
  INSERT INTO users( username, email, name, surname, password, user_type)
	VALUES ('yahyabasakci', 'yahya@deneme.com', 'yahya', 'basakci', '123456','admin');
  INSERT INTO users( username, email, name, surname, password, user_type)
	VALUES ('nazararik', 'nazar@deneme.com', 'nazar', 'arik', '123456','admin');

`

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

	if err := Db.AutoMigrate(&models.User{}, &models.Answer{}, &models.Quiz{}, &models.Question{}, &models.Lesson{}); err != nil {
		panic("failed to migrate database")
	}

	Db.Exec(createSQL)

	fmt.Println("Connection Opened to Database")

}
