package routes

import (
	controllers "dealsafe/controllers/auth"

	"github.com/gin-gonic/gin"
)


func SetUpRoutes(route *gin.Engine){

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// auth
	route.POST("/register", controllers.RegisterUser)
	route.POST("/signin", controllers.SignIn)
	route.GET("/profile", controllers.GetProfile)


}