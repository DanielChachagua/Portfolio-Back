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

func CreateProject(image *multipart.FileHeader, project *models.CreateProject) (string, error) {
	name, err := utils.SaveImage(image)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al guardar la imagen", err)
	}

	projectID, err := repositories.Repo.CreateProject(name, project)

	if err != nil {
		utils.DeleteImage(name)
		return "", models.ErrorResponse(500, "Error al crear el proyecto", err)
	}

	return projectID, nil
}

func GetProjectByID(id string, baseUrl string) (*models.Project, error) {
	project, err := repositories.Repo.GetProjectByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar el proyecto", err)
	}

	(*project).UrlImage = fmt.Sprintf("%s/image/get/%s", baseUrl, project.UrlImage)
	return repositories.Repo.GetProjectByID(id)
}

func GetAllProjects(baseUrl string) (*[]models.Project, error) {
	projects, err := repositories.Repo.GetAllProjects()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener los proyectos", err)
	}

	for i, project := range *projects {
		(*projects)[i].UrlImage = fmt.Sprintf("%s/image/get/%s", baseUrl, project.UrlImage)
	}
	return projects, nil
}

func GetFavorites(baseUrl string) (*[]models.Project, error) {
	projects, err := repositories.Repo.GetFavorites()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener los proyectos", err)
	}

	for i, project := range *projects {
		(*projects)[i].UrlImage = fmt.Sprintf("%s/image/get/%s", baseUrl, project.UrlImage)
	}

	return projects, nil
}

func UpdateProject(id string, image *multipart.FileHeader, projectUpdate *models.UpdateProject) error {
	var urlImage string
	if image != nil {
		name, err := utils.SaveImage(image)
		if err != nil {
			return models.ErrorResponse(500, "Error al guardar la imagen", err)
		}
		urlImage = name
	}
	err := repositories.Repo.UpdateProject(id, urlImage, projectUpdate)
	if err != nil {
		if urlImage != "" {
			utils.DeleteImage(urlImage)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar el proyecto", err)
	}

	return nil
}

func DeleteProject(id string) error {
	err := repositories.Repo.DeleteProject(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Proyecto no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al buscar el proyecto", err)
	}

	return nil
}