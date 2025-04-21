package repositories

import (
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/google/uuid"
)

func (r *Repository) GetProjectByID(id string) (*models.Project, error) {
	var project models.Project
	err := r.DB.First(&project, id).Error
	
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *Repository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Find(&projects).Error
	
	if err != nil {
		return nil, err
	}

	return projects, nil
}
func (r *Repository) UpdateProject(id string, urlImage string, projectUpdate *models.UpdateProject) (string, error) {
	var project models.Project
	err := r.DB.First(&project, id).Error
	if err != nil {
		return "", err
	}

	project.Title = projectUpdate.Title
	project.Description = projectUpdate.Description
	if urlImage != "" {
		project.UrlImage = urlImage
	}
	project.Link = projectUpdate.Link
	project.UpdatedAt = time.Now().UTC()

	err = r.DB.Save(&project).Error
	if err != nil {
		return "", err
	}

	return project.ID, nil
}
func (r *Repository) DeleteProject(id string) error {
	var project models.Project
	err := r.DB.First(&project, id).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&project).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateProject(urlImage string, project *models.CreateProject) (string, error) {
	err := r.DB.Create(&project).Error
	if err != nil {
		return "", err
	}

	newId := uuid.NewString()
	r.DB.Create(&models.Project{
		ID:          newId,
		Title:       project.Title,
		Description: project.Description,
		UrlImage:    urlImage,
		Link:        project.Link,
	})

	return newId, nil
}