package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:30" json:"first_name"`
	LastName  string    `gorm:"size:30" json:"last_name"`
	Username  string    `gorm:"unique;size:30;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	UrlImage  string    `gorm:"type:text" json:"url_image"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserCreate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	UrlImage  string `json:"url_image" validate:"required"`
}

func (u *UserCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `son:"last_name"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	UrlImage  string `json:"url_image" validate:"required"`
}

func (u *UserUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserResponse struct {
	ID       string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `son:"last_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UrlImage string `json:"url_image"`
}

type PasswordUpdate struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (p *PasswordUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
