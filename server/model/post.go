package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null"`
	Content  string    `json:"content" gorm:"not null"`
	UserId   uint      `json:"user_id" gorm:"<-:create"`
	Comments []Comment `json:"comments"`
}
