package controllers

import (
	"dealsafe/database"
	"dealsafe/lib"
	"github.com/gin-gonic/gin"
)



func GetProfile(c *gin.Context) {
	// Get the token from cookies
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized: No token provided"})
		return
	}

	claims, err := lib.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	user, err := database.Queries.FindUserByEmail(c, claims["email"].(string))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile retrieved successfully",
		"user": gin.H{
			"user": user,
			// You can query the database here using database.Queries.FindUserByID(c, uuid)
		},
	})
}