package services

import (
	"errors"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"gorm.io/gorm"
)

func AuthLogin(email string, password string) (string, error) {
	user, err := repositories.Repo.GetUserByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al  buscar usuario", err)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", models.ErrorResponse(401, "Credenciales incorrectas", nil)
	}

	token, err := utils.GenerateUserToken(user)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar token", err)
	}

	return token, nil
}

func CurrentUser(id string) (*models.User, error) {
	user, err := repositories.Repo.GetUserByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}

	return user, nil
}