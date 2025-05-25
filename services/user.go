package services

import (
	"errors"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"gorm.io/gorm"
)

func CreateUser(user *models.UserCreate) (string, error) {
	return repositories.Repo.CreateUser(user)
}

func GetUser() (*models.UserResponse, error) {
	user, err := repositories.Repo.GetUser()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ErrorStruc{
				StatusCode: 404,
				Message:    "User not found",
			}
		}
		return nil, err
	}
	return user, nil
}