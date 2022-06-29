package controllers

import (
	"BootcampHacktiv8/final_project/helpers"
	"BootcampHacktiv8/final_project/models"
	"BootcampHacktiv8/final_project/repositories"
	"BootcampHacktiv8/final_project/validations"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password with err: %s", err)
		helpers.NewHandlerResponse("Error hash", nil).Failed(c)
		return
	}
	user.Password = string(password)

	err = repositories.RegisterRepository(&user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully register", nil).SuccessCreate(c)
	// _, err := govalidator.ValidateStruct(&user)
	// if err != nil {
	// errs := err.(govalidator.Errors).Errors()
	// fmt.Println(errs)
	// for _, e := range errs {
	// 	fmt.Println(e.Error())
	// }
	// fmt.Println("masuk")
	// 	helpers.NewHandlerValidationResponse(err.Error(), nil).BadRequest(c)
	// 	return
	// }
}

func LoginController(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}
	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	userData, err := repositories.LoginRepository(&user)
	if err != nil {
		helpers.NewHandlerResponse("Email not found", nil).Failed(c)
		return
	}
	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if errHash != nil {
		log.Printf("Email or Password Incorrect with err: %s\n", errHash)
		helpers.NewHandlerResponse("Email or password is incorrect", nil).BadRequest(c)
		return
	}
	tokenString, err := helpers.GenerateJWT(userData)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		c.Abort()
		return
	}
	id := strconv.Itoa(userData.Id)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24 * 60),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "id",
		Value:   id,
		Expires: time.Now().Add(time.Hour * 24 * 60),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	helpers.NewHandlerResponse("Successfully Login", nil).Success(c)
}

func LogoutController(c *gin.Context) {
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
	helpers.NewHandlerResponse("Successfully logout", nil).Success(c)
}
