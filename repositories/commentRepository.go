package repositories

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/models"
	"log"
)

func CreateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error) {
	db := db.DbManager()
	comment.UserId = authId
	err := db.Create(&comment)
	if err.Error != nil {
		return nil, err.Error
	}
	return comment, nil
}

func GetAllCommentRepository() (*[]models.CommentResponse, error) {
	db := db.DbManager()
	comments := []models.CommentResponse{}
	if err := db.Joins("User").Joins("Photo").Model(&models.Comment{}).Find(&comments).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &comments, nil
}

func UpdateCommentRepository(authId int, comment *models.Comment) (*models.Comment, error) {
	db := db.DbManager()
	commentCheckAvail := models.Comment{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", comment.Id).
		First(&commentCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data comment detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = db.Model(&commentCheckAvail).Updates(comment)
	if err.Error != nil {
		log.Printf("Error update data comment id with err: %s", err.Error)
		return nil, err.Error
	}
	return &commentCheckAvail, nil
}

func DeleteCommentRepository(commentId, authId int) error {
	db := db.DbManager()
	commentCheckAvail := models.Comment{}
	err := db.Where("user_id = ?", authId).
		Where("id = ?", commentId).
		First(&commentCheckAvail)
	if err.Error != nil {
		log.Printf("Error delete comment detail with err: %s", err.Error)
		return err.Error
	}
	db.Delete(commentCheckAvail, commentId)
	return nil
}
