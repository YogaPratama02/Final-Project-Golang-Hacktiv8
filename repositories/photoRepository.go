package repositories

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/models"
	"log"
)

func CreatePhotoRepository(Id int, photo *models.Photo) (*models.Photo, error) {
	db := db.DbManager()
	photo.UserId = Id
	err := db.Create(&photo)
	if err.Error != nil {
		return nil, err.Error
	}
	return photo, nil
}

func GetAllPhotoRepository() (*[]models.PhotoResponse, error) {
	db := db.DbManager()
	photos := []models.PhotoResponse{}
	if err := db.Joins("User").Model(&models.Photo{}).Find(&photos).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &photos, nil
}

func UpdatePhotoRepository(authId int, photo *models.Photo) (*models.Photo, error) {
	db := db.DbManager()
	photoCheckAvail := models.Photo{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", photo.Id).
		First(&photoCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data photo detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = db.Model(&photoCheckAvail).Updates(photo)
	if err.Error != nil {
		log.Printf("Error update data photo id with err: %s", err.Error)
		return nil, err.Error
	}
	return &photoCheckAvail, nil
}

func DeletePhotoRepository(photoId, authId int) error {
	db := db.DbManager()
	photoCheckAvail := models.Photo{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", photoId).
		First(&photoCheckAvail)
	if err.Error != nil {
		log.Printf("Error delete photo detail with err: %s", err.Error)
		return err.Error
	}
	db.Delete(photoCheckAvail, photoId)
	return nil
}
