package routes

import (
	"BootcampHacktiv8/final_project/controllers"
	"BootcampHacktiv8/final_project/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	// binding.Validator = new(validations.DefaultValidator)
	e := gin.New()
	e.Use(gin.Recovery(), middlewares.Logger())
	v1 := e.Group("/api/v1")
	user := v1.Group("/users")
	{
		user.POST("/register", controllers.RegisterController)
		user.POST("/login", controllers.LoginController)
	}
	secured := v1.Group("/").Use(middlewares.Auth())
	{
		secured.PUT("/users/:id", controllers.UpdateUserController)
		secured.DELETE("/users", controllers.DeleteUserController)
		secured.POST("users/logout", controllers.LogoutController)
	}
	photo := v1.Group("photos/").Use(middlewares.Auth())
	{
		photo.POST("", controllers.CreatePhotoController)
		photo.GET("", controllers.GetAllPhotoController)
		photo.PUT(":photoId", controllers.UpdatePhotoController)
		photo.DELETE(":photoId", controllers.DeletePhotoController)
	}

	comment := v1.Group("comment/").Use(middlewares.Auth())
	{
		comment.POST("", controllers.CreateCommentController)
		comment.GET("", controllers.GetAllCommentController)
		comment.PUT(":commentId", controllers.UpdateCommentController)
		comment.DELETE(":commentId", controllers.DeleteCommentController)
	}

	socialMedia := v1.Group("socialmedias").Use(middlewares.Auth())
	{
		socialMedia.POST("", controllers.CreateSocialMediaController)
		socialMedia.GET("", controllers.GetAllSocialMediaController)
		socialMedia.PUT(":socialMediaId", controllers.UpdateSocialMediaController)
		socialMedia.DELETE(":socialMediaId", controllers.DeleteSocialMediaController)
	}
	e.Run(":8000")
}
