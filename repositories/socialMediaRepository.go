package repositories

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/models"
	"log"
)

func CreateSocialMediaRepository(authId int, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	db := db.DbManager()
	socialMedia.UserId = authId
	err := db.Create(&socialMedia)
	if err.Error != nil {
		return nil, err.Error
	}
	return socialMedia, nil
}

func GetAllSocialMediaRepository() (*[]models.SocialMediaResponse, error) {
	db := db.DbManager()
	socialMedia := []models.SocialMediaResponse{}
	if err := db.Joins("User").Model(&models.SocialMedia{}).Find(&socialMedia).Error; err != nil {
		log.Printf("Error get data social media with err: %s", err)
		return nil, err
	}
	return &socialMedia, nil
}

func UpdateSocialMediaRepository(authId int, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	db := db.DbManager()
	socialMediaCheckAvail := models.SocialMedia{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", socialMedia.Id).
		First(&socialMediaCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data social media detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = db.Model(&socialMediaCheckAvail).Updates(socialMedia)
	if err.Error != nil {
		log.Printf("Error update data social media id with err: %s", err.Error)
		return nil, err.Error
	}
	return &socialMediaCheckAvail, nil
}

func DeleteSocialMediaRepository(socialMediaId, authId int) error {
	db := db.DbManager()
	socialMediaCheckAvail := models.SocialMedia{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", socialMediaId).
		First(&socialMediaCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data social media detail with err: %s", err.Error)
		return err.Error
	}
	db.Delete(socialMediaCheckAvail, socialMediaId)
	return nil
}
