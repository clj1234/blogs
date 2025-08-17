package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"type:varchar(255);uniqueIndex"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Posts    []Post    `json:"posts"`
	Comments []Comment `json:"comments"`
}
