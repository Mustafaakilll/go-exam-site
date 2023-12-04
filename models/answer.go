package models

import "time"

type Answer struct {
	AnswerID   int      `gorm:"primaryKey;autoIncrement" json:"answer_id"`
	QuestionID int      `gorm:"not null" json:"question_id"`
	AnswerText string   `gorm:"type:text;not null" json:"answer_text"`
	IsCorrect  bool     `gorm:"not null" json:"is_correct"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"question"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
