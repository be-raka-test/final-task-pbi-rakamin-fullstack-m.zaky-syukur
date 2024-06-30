package routes

import (
	"btpn-go/controllers"
	"btpn-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/users", controllers.GetUsers)

		api.Use(middlewares.AuthMiddleware())
		{
			api.GET("/photos", controllers.GetPhotos)
			api.GET("/photos/:id", controllers.GetPhoto)
			api.POST("/photos", controllers.CreatePhoto)
			api.PUT("/photos/:id", controllers.UpdatePhoto)
			api.DELETE("/photos/:id", controllers.DeletePhoto)
		}
	}
}
