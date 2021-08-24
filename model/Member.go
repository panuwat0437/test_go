package model

import "time"

// https://gorm.io/docs/models.html

type Member struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Email string `form:"email" binding:"required"`
	Level     string `gorm:"default:normal"`
	CreatedAt time.Time
}
