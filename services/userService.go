package services

import (
	"SIMPLE-API/config"
	"SIMPLE-API/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func CreateUser(user models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

func UpdateUser(id int, updatedUser models.User) error {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = config.DB.Model(&user).Updates(updatedUser)
	return result.Error
}

func DeleteUser(id int) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
