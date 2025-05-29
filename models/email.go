package models

import "github.com/go-playground/validator/v10"

type EmailContact struct {
	Issue string `json:"issue" validate:"required"`
	Body  string `json:"body" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func (e *EmailContact) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}