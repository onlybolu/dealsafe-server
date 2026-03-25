package main

import (
	"dealsafe/database"
	"dealsafe/lib"
	"dealsafe/routes"
	"fmt"
	"dealsafe/middlewares"
	"github.com/gin-gonic/gin"
)

func init(){
	lib.LoadEnv()
}

func main(){

	router := gin.Default()

	// middlewares
	router.Use(middlewares.RateLimit())
	
	database.ConnectDB()

	routes.SetUpRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}