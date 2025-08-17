package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" `
	UserId  uint   `json:"user_id" gorm:"not null" gorm:"<-:create"`
	PostId  uint   `json:"post_id" gorm:"not null" gorm:"<-:create"`
}
