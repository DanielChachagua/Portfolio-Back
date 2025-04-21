package models

import (
	"time"
)

type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"unique;size:30;not null" json:"username"`
	Email     string     `gorm:"unique;not null" json:"email"`         
	Password  string     `gorm:"not null" json:"password"`             
	UrlImage  *string    `gorm:"type:text" json:"url_image"`               
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`      
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`        
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UrlImage *string `json:"url_image"`
}
