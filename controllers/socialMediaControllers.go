package controllers

import (
	"BootcampHacktiv8/final_project/helpers"
	"BootcampHacktiv8/final_project/models"
	"BootcampHacktiv8/final_project/repositories"
	"BootcampHacktiv8/final_project/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSocialMediaController(c *gin.Context) {
	var socialMedia models.SocialMedia
	authId := helpers.ParseJWT(c)
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&socialMedia); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	data, err := repositories.CreateSocialMediaRepository(authId, &socialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create social media", data).SuccessCreate(c)
}

func GetAllSocialMediaController(c *gin.Context) {
	data, err := repositories.GetAllSocialMediaRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all social media", data).Success(c)
}

func UpdateSocialMediaController(c *gin.Context) {
	var socialMedia models.SocialMedia

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	socialMedia.Id = socialMediaId
	authId := helpers.ParseJWT(c)
	if err := validations.DoValidation(&socialMedia); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	data, err := repositories.UpdateSocialMediaRepository(authId, &socialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update social media", data).Success(c)
}

func DeleteSocialMediaController(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	authId := helpers.ParseJWT(c)
	err = repositories.DeleteSocialMediaRepository(socialMediaId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your social media has been successfully delete", nil).Success(c)
}
