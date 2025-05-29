package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetProjectByID(id string) (*models.Project, error) {
	var project models.Project
	if err := r.DB.Preload("Skills").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *Repository) GetAllProjects() (*[]models.Project, error) {
	var projects []models.Project
	err := r.DB.Preload("Skills").Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *Repository) GetFavorites() (*[]models.Project, error) {
	var projects []models.Project
	err := r.DB.Preload("Skills").Find(&projects).Where("favorite = ?", true).Error

	if err != nil {
		return nil, err
	}

	return &projects, nil
}

// func (r *Repository) UpdateProject(id string, urlImage string, projectUpdate *models.UpdateProject) error {
// 	var project models.Project
// 	err := r.DB.First(&project, id).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return models.ErrorResponse(404, "Proyecto no encontrado", err)
// 		}
// 		return err
// 	}

// 	project.Title = projectUpdate.Title
// 	project.Description = projectUpdate.Description
// 	if urlImage != "" {
// 		project.UrlImage = urlImage
// 	}
// 	project.Link = projectUpdate.Link
// 	project.Favorite = projectUpdate.Favorite
// 	project.UpdatedAt = time.Now().UTC()

// 	err = r.DB.Save(&project).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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

	err = utils.DeleteImage(project.UrlImage)
	if err != nil {
		return models.ErrorResponse(500, "Error al eliminar la imagen del proyecto", err)
	}

	return nil
}

// func (r *Repository) CreateProject(urlImage string, project *models.CreateProject) (string, error) {
// 	newId := uuid.NewString()

// 	if err := r.DB.Create(&models.Project{
// 		ID:          newId,
// 		Title:       project.Title,
// 		Description: project.Description,
// 		UrlImage:    urlImage,
// 		Link:        project.Link,
// 		Favorite:    project.Favorite,
// 	}).Error; err != nil {
// 		return "", err
// 	}

// 	return newId, nil
// }

func (r *Repository) CreateProject(urlImage string, projectData *models.CreateProject,) (string, error) {
	var skills []models.Skill
	if err := r.DB.Where("id IN ?", projectData.SkillsID).Find(&skills).Error; err != nil {
		return "", err
	}

	project := models.Project{
		ID:          uuid.NewString(),
		Title:       projectData.Title,
		Description: projectData.Description,
		UrlImage:    urlImage,
		Link:        projectData.Link,
		Favorite:    projectData.Favorite,
		Skills:      skills,
	}

	if err := r.DB.Create(&project).Error; err != nil {
		return "", err
	}

	return project.ID, nil
}

func (r *Repository) UpdateProject(id string, urlImage string, projectUpdate *models.UpdateProject) error {
	var project models.Project
	if err := r.DB.First(&project, "id = ?", id).Error; err != nil {
		return err
	}

	project.Title = projectUpdate.Title
	project.Description = projectUpdate.Description
	project.Link = projectUpdate.Link
	project.Favorite = projectUpdate.Favorite
	if urlImage != "" {
		err := utils.DeleteImage(project.UrlImage)
		if err != nil {
			return err
		}
		project.UrlImage = urlImage
	}
	project.UpdatedAt = time.Now().UTC()

	var skills []models.Skill
	if err := r.DB.Where("id IN ?", projectUpdate.SkillsID).Find(&skills).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&project).Association("Skills").Replace(skills); err != nil {
		return err
	}

	if err := r.DB.Save(&project).Error; err != nil {
		return err
	}

	return nil
}
