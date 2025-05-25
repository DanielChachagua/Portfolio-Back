package repositories

import (
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/google/uuid"
)

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateUser(id string, userUpdate *models.UserUpdate) (string, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return "", err
	}

	user.UpdatedAt = time.Now().UTC()
	user.Username = userUpdate.Username
	user.Email = userUpdate.Email
	user.UrlImage = userUpdate.UrlImage

	err = r.DB.Save(&user).Error
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *Repository) CreateUser(user *models.UserCreate) (string, error) {
	id := uuid.New().String()
	newUser := models.User{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		UrlImage: user.UrlImage,
	}

	err := r.DB.Create(&newUser).Error
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *Repository) GetUser() (*models.UserResponse, error) {
	var user models.User
	err := r.DB.First(&user).Error
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		UrlImage: user.UrlImage,
	}

	return userResponse, nil
}