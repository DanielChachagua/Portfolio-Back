package services

import (
	"mime/multipart"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
)

func CreateProject(image *multipart.FileHeader, project *models.CreateProject) (string, error) {
	name, err := SaveImage(image)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al guardar la imagen", err)
	}

	projectID, err := repositories.Repo.CreateProject(name, project)

	if err != nil {
		DeleteImage(name)
		return "", models.ErrorResponse(500, "Error al crear el proyecto", err)
	}

	return projectID, nil
}