package models

import "time"

type Quiz struct {
	QuizID      int    `gorm:"primaryKey;autoIncrement" json:"quiz_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	TeacherID   int    `gorm:"not null" json:"teacher_id"`
	Teacher     User   `gorm:"foreignKey:TeacherID" json:"teacher"`

	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
