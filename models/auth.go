package models

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	esTranslations "github.com/go-playground/validator/v10/translations/es"
)

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (l *Login) Validate() error {
	validate := validator.New()
	translator := es.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("es")
	esTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(l)
	if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			var msg string
			for _, e := range validationErrors {
					switch e.Field() {
					case "Email":
							if e.Tag() == "required" {
									msg = "El correo es obligatorio"
							} else if e.Tag() == "email" {
									msg = "El correo no es válido"
							}
					case "Password":
							if e.Tag() == "required" {
									msg = "La contraseña es obligatoria"
							} else if e.Tag() == "min" {
									msg = "La contraseña debe tener al menos 6 caracteres"
							}
					}
					if msg != "" {
						return errors.New(msg)
					}
			}
	}
	return nil
}