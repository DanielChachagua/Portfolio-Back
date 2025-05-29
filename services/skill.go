package services

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"gorm.io/gorm"
)

func SkillCreate(skill *models.SkillCreate, file *multipart.FileHeader) (string, error) {
	nameImage, err := utils.SaveImage(file)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al guardar la imagen", err)
	}

	id, err := repositories.Repo.CreateSkill(skill, nameImage)
	if err != nil {
		utils.DeleteImage(nameImage)
		return "", models.ErrorResponse(500, "Error al guardar la skill", err)
	}

	return id, nil
}

func SkillGetAll(baseUrl string) (*[]models.Skill, error) {
	skills, err := repositories.Repo.GetAllSkill()
	if err != nil {
		return nil, err
	}

	for i, skill := range *skills {
		(*skills)[i].UrlImage = fmt.Sprintf("%s/image/get/%s", baseUrl, skill.UrlImage)
	}

	return skills, nil
}

func SkillUpdate(id string, skill *models.SkillUpdate, file *multipart.FileHeader) error {
	fileName, err := utils.SaveImage(file)
	if err != nil {
		return models.ErrorResponse(500, "Error al guardar la imagen", err)
	}

	if err := repositories.Repo.UpdateSkill(id, skill, fileName); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if fileName != "" {
				utils.DeleteImage(fileName)
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Skill no encontrado", err)
			}
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