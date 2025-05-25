package models

import (
	"errors"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	esTranslations "github.com/go-playground/validator/v10/translations/es"
	"time"
)

type Project struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Link        string   `gorm:"type:text" json:"link"`
	UrlImage    string    `gorm:"type:text" json:"image"`
	Favorite    bool      `gorm:"default:false" json:"favorite"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type CreateProject struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Link        string `json:"link"`
	Favorite    bool     `json:"favorite" validate:"default=false"`
}

func (cp *CreateProject) Validate() error {
	validate := validator.New()
	translator := es.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("es")
	esTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(cp)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var msg string
		for _, e := range validationErrors {
			switch e.Field() {
				case "Title":
					if e.Tag() == "required" {
						msg = "El titulo es obligatorio"
					}
				case "Description":
					if e.Tag() == "required" {
						msg = "La descripción es obligatoria"
					}
			}
			if msg != "" {
				return errors.New(msg)
			}
		}
	}
	return nil
}

type UpdateProject struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	Link        string `json:"link" validate:"required"`
	Favorite    bool     `json:"favorite" validate:"default=false"`
}

func (up *UpdateProject) Validate() error {
	validate := validator.New()
	translator := es.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("es")
	esTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(up)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var msg string
		for _, e := range validationErrors {
			switch e.Field() {
				case "Title":
					if e.Tag() == "required" {
						msg = "El titulo es obligatorio"
					}
				case "Description":
					if e.Tag() == "required" {
						msg = "La descripción es obligatoria"
					}
				case "Link":
					if e.Tag() == "required" {
						msg = "El link es obligatorio"
					}
			}
			if msg != "" {
				return errors.New(msg)
			}
		}
	}
	return nil
}
