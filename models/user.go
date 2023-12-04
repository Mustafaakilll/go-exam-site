package models

import "time"

type UserType string

const (
	Admin   UserType = "admin"
	Teacher UserType = "teacher"
	Student UserType = "student"
)

func (r *UserType) Scan(value any) {
	*r = UserType(value.([]byte))
}

type User struct {
	UserID   int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"not null" json:"email"`
	UserType string `gorm:"type:UserType;not null" json:"user_type"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
