package repositories

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/google/uuid"
)

func (r *Repository) CreateSkill(skill *models.SkillCreate) (string, error) {
	uuid := uuid.NewString()

	if err := r.DB.Create(&models.Skill{
		ID:   uuid,
		Name: skill.Name,
	}).Error; err != nil {
		return "", err
	}

	return uuid, nil
}

func (r *Repository) GetAllSkill() (*[]models.Skill, error) {
	var skills []models.Skill

	if err := r.DB.Find(&skills).Error; err != nil {
		return nil, err
	}

	return &skills, nil
}

func (r *Repository) UpdateSkill(id string, skill *models.SkillUpdate) error {
	if err := r.DB.Model(&models.Skill{}).Where("id = ?", id).Updates(models.Skill{
		Name: skill.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSkill(id string) error {
	if err := r.DB.Delete(&models.Skill{}, id).Error; err != nil {
		return err
	}

	return nil
}