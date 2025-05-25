package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetProjectByID(id string) (*models.Project, error) {
	var project models.Project
	err := r.DB.First(&project, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return nil, err
	}

	return &project, nil
}

func (r *Repository) GetAllProjects() (*[]models.Project, error) {
	var projects []models.Project
	err := r.DB.Find(&projects).Error
	
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *Repository) GetFavorites() (*[]models.Project, error) {
	var projects []models.Project
	err := r.DB.Find(&projects).Where("favorite = ?", true).Error
	
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *Repository) UpdateProject(id string, urlImage string, projectUpdate *models.UpdateProject) error {
	var project models.Project
	err := r.DB.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return err
	}

	project.Title = projectUpdate.Title
	project.Description = projectUpdate.Description
	if urlImage != "" {
		project.UrlImage = urlImage
	}
	project.Link = projectUpdate.Link
	project.Favorite = projectUpdate.Favorite
	project.UpdatedAt = time.Now().UTC()

	err = r.DB.Save(&project).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteProject(id string) error {
	var project models.Project
	err := r.DB.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return err
	}

	err = r.DB.Delete(&project).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateProject(urlImage string, project *models.CreateProject) (string, error) {
	newId := uuid.NewString()

	if err := r.DB.Create(&models.Project{
		ID:          newId,
		Title:       project.Title,
		Description: project.Description,
		UrlImage:    urlImage,
		Link:        project.Link,
		Favorite: 	project.Favorite,
	}).Error; err != nil {
		return "", err
	}

	return newId, nil
}
