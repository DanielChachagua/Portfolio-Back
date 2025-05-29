package models

import "github.com/go-playground/validator/v10"

type Skill struct {
	ID   string    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
	Area string `gorm:"size:100;not null" json:"area" validate:"required,oneof=back front tool"`
	UrlImage string `gorm:"type:text" json:"url_image"`
	Projects []Project `gorm:"many2many:project_skills;" json:"projects"`
}

type SkillCreate struct {
	Name string `json:"name" validate:"required"`
	Area string `json:"area" validate:"required,oneof=back front tool"`
}

func (s *SkillCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type SkillUpdate struct {
	Name string `json:"name" validate:"required"`
	Area string `json:"area" validate:"required,oneof=back front tool"`
}

func (s *SkillUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}