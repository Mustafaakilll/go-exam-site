package models

type Lesson struct {
	LessonID   int    `gorm:"primaryKey;autoIncrement" json:"lesson_id"`
	LessonName string `gorm:"not null" json:"lesson_name"`
}
