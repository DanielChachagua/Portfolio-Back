package repositories

import (
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
)

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(emial string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", emial).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateUser(id string, userUpdate models.UserUpdate) (string, error) {
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