package controllers

import (
	"BootcampHacktiv8/final_project/helpers"
	"BootcampHacktiv8/final_project/models"
	"BootcampHacktiv8/final_project/repositories"
	"BootcampHacktiv8/final_project/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhotoController(c *gin.Context) {
	var photo models.Photo
	authId := helpers.ParseJWT(c)
	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&photo); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	data, err := repositories.CreatePhotoRepository(authId, &photo)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create photo", data).SuccessCreate(c)
}

func GetAllPhotoController(c *gin.Context) {
	data, err := repositories.GetAllPhotoRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all photo", data).Success(c)
}

func UpdatePhotoController(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	photo.Id = photoId
	authId := helpers.ParseJWT(c)
	if err := validations.DoValidation(&photo); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	data, err := repositories.UpdatePhotoRepository(authId, &photo)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update photo", data).Success(c)
}

func DeletePhotoController(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	authId := helpers.ParseJWT(c)
	err = repositories.DeletePhotoRepository(photoId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your photo has been successfully delete", nil).Success(c)
}
