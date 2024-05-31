package router

import (
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/controllers"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.POST("/logout", controllers.Logout)
		userRoutes.PUT("/:userId", middleware.CheckAuth, controllers.UpdateUser)
		userRoutes.DELETE("/:userId", middleware.CheckAuth, controllers.DeleteUser)
	}

	// Photo
	photoRoutes := r.Group("/photos")
	photoRoutes.Use(middleware.CheckAuth)
	{
		photoRoutes.GET("/", controllers.ShowPhoto)
		photoRoutes.POST("/", controllers.CreatePhoto)
		photoRoutes.PUT("/:photoId", controllers.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r
}
