package models

import "time"

type QuestionType string

const (
	True_False      QuestionType = "True/False"
	Multiple_Choice QuestionType = "Multiple Choice"
	Open_Ended      QuestionType = "Open-Ended"
)

func (qt *QuestionType) Scan(value any) {
	*qt = QuestionType(value.([]byte))
}

type Question struct {
	QuestionID   int    `gorm:"primaryKey;autoIncrement" json:"question_id"`
	QuizID       int    `gorm:"not null" json:"quiz_id"`
	QuestionText string `gorm:"type:text;not null" json:"question_text"`
	QuestionType int    `gorm:"not null" json:"question_type"`
	Quiz         Quiz   `gorm:"foreignKey:QuizID" json:"quiz"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
