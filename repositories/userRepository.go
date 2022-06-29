package repositories

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/models"
	"log"
)

func UpdateUserRepository(Id int, user *models.UserUpdate) (*models.User, error) {
	db := db.DbManager()
	userCheckId := models.User{}
	err := db.Select("id", "email", "user_name", "age", "created_at", "updated_at").First(&userCheckId, Id)
	if err.Error != nil {
		log.Printf("Error get data user id with err: %s", err.Error)
		return nil, err.Error
	}

	err = db.Model(&userCheckId).Updates(models.User{Email: user.Email, UserName: user.UserName})
	if err.Error != nil {
		log.Printf("Error update data user id with err: %s", err.Error)
		return nil, err.Error
	}

	return &userCheckId, nil
}

func DeleteUserRepository(Id int) error {
	db := db.DbManager()
	userCheckId := models.User{}
	err := db.Select("id", "email", "user_name", "age", "created_at", "updated_at").First(&userCheckId, Id)
	if err.Error != nil {
		log.Printf("Error get data user id with err: %s", err.Error)
		return err.Error
	}
	db.Delete(userCheckId, Id)
	return nil
}
