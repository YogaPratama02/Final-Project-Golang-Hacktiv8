package controllers

import (
	"BootcampHacktiv8/final_project/helpers"
	"BootcampHacktiv8/final_project/models"
	"BootcampHacktiv8/final_project/repositories"
	"BootcampHacktiv8/final_project/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommentController(c *gin.Context) {
	var comment models.Comment
	authId := helpers.ParseJWT(c)
	if err := c.ShouldBindJSON(&comment); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&comment); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	data, err := repositories.CreateCommentRepository(authId, &comment)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create comment", data).SuccessCreate(c)
}

func GetAllCommentController(c *gin.Context) {
	data, err := repositories.GetAllCommentRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all comment", data).Success(c)
}

func UpdateCommentController(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	comment.Id = commentId
	authId := helpers.ParseJWT(c)
	if err := validations.DoValidation(&comment); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	data, err := repositories.UpdateCommentRepository(authId, &comment)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update comment", data).Success(c)
}

func DeleteCommentController(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	authId := helpers.ParseJWT(c)
	err = repositories.DeleteCommentRepository(commentId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your comment has been successfully delete", nil).Success(c)
}
