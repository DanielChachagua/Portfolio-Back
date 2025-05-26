package services

import (
	"errors"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"gorm.io/gorm"
)

func SkillCreate(skill *models.SkillCreate) (string, error) {
	if err := skill.Validate(); err != nil {
		return "", err
	}

	id, err := repositories.Repo.CreateSkill(skill)
	if err != nil {
		return "", err
	}

	return id, nil
}

func SkillGetAll() (*[]models.Skill, error) {
	skills, err := repositories.Repo.GetAllSkill()
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func SkillUpdate(id string, skill *models.SkillUpdate) error {
	if err := repositories.Repo.UpdateSkill(id, skill); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("skill not found")
		}
		return err
	}

	return nil
}

func SkillDelete(id string) error {
	if err := repositories.Repo.DeleteSkill(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("skill not found")
		}
		return err
	}

	return nil
}	