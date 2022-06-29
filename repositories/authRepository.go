package repositories

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/models"
	"log"
)

func RegisterRepository(user *models.User) error {
	db := db.DbManager()
	err := db.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func LoginRepository(user *models.UserLogin) (*models.UserLogin, error) {
	db := db.DbManager()
	var userLogin models.UserLogin
	if err := db.Table("users").Where("users.email = ?", user.Email).First(&userLogin).Error; err != nil {
		log.Printf("Error get data user with err: %s", err)
		return nil, err
	}
	return &userLogin, nil
}
