package router

import (
	"github.com/zakkaizzatur/golang-dts-final-project/controllers"
	"github.com/zakkaizzatur/golang-dts-final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

		userRouter.PUT("/:userId",middlewares.Authentication(), middlewares.UserAuthorization(), controllers.UpdateUser)

		userRouter.DELETE("/:userId",middlewares.Authentication(), middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())

		photosRouter.POST("/", controllers.CreatePhoto)

		photosRouter.GET("/", controllers.GetPhotos)

		photosRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		photosRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)

	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())

		commentsRouter.POST("/", controllers.CreateComment)

		commentsRouter.GET("/", controllers.GetComments)

		commentsRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
	
		commentsRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())

		socialMediaRouter.POST("/", controllers.CreateSocialMedia)

		socialMediaRouter.GET("/", controllers.GetSocialMedia)

		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
	
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}