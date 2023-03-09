package routes

import (
	"github.com/gabhmm/rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBook)
			books.POST("/:id", controllers.ShowBooks)

		}
	}
	return router
}
