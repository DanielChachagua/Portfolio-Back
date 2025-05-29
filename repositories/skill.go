package repositories

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"github.com/google/uuid"
)

func (r *Repository) CreateSkill(skill *models.SkillCreate, nameImage string) (string, error) {
	uuid := uuid.NewString()

	if err := r.DB.Create(&models.Skill{
		ID:   uuid,
		Name: skill.Name,
		Area: skill.Area,
		UrlImage: nameImage,
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

func (r *Repository) UpdateSkill(id string, skillUpdate *models.SkillUpdate, urlImage string) error {
	var skill models.Skill
	if err := r.DB.First(&skill, "id = ?", id).Error; err != nil {
		return err
	}

	skill.Name = skillUpdate.Name
	skill.Area = skillUpdate.Area
	if urlImage != "" {
		err := utils.DeleteImage(skill.UrlImage)
		if err != nil {
			return err
		}
		skill.UrlImage = urlImage
	}
	
	if err := r.DB.Model(&models.Skill{}).Where("id = ?", id).Updates(models.Skill{
		Name: skill.Name,
		Area: skill.Area,
		
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSkill(id string) error {
	var skill models.Skill
	if err := r.DB.First(&skill, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.DB.Delete(&skill).Error; err != nil {
		return err
	}
	
	err := utils.DeleteImage(skill.UrlImage)
	if err != nil {
		return err
	}

	return nil
}