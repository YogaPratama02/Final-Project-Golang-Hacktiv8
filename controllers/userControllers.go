package controllers

import (
	"BootcampHacktiv8/final_project/helpers"
	"BootcampHacktiv8/final_project/models"
	"BootcampHacktiv8/final_project/repositories"
	"BootcampHacktiv8/final_project/validations"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateUserController(c *gin.Context) {
	var user models.UserUpdate
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}
	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	data, err := repositories.UpdateUserRepository(id, &user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update user", data).Success(c)
}

func DeleteUserController(c *gin.Context) {
	authId := helpers.ParseJWT(c)
	err := repositories.DeleteUserRepository(authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "id",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	helpers.NewHandlerResponse("Your account has been successfully deleted", nil).Success(c)
}
