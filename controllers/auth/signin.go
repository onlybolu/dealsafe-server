package controllers

import (
	"dealsafe/database"
	"dealsafe/lib"
	request "dealsafe/types"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func SignIn(c *gin.Context) {
	var authRequest request.AuthRequest

	if err := c.BindJSON(&authRequest); err != nil{
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error": err.Error(),
		})
		return
	}

	existingUser, err := database.Queries.FindUserByEmail(
		c,
		authRequest.Email,
	)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "user does not exist",
			"error": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(authRequest.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Password incorrect! please enter the correct password",
		})
		return
	}


	access_token, err := lib.GenerateJWT(existingUser.Email, existingUser.ID.String())

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to generate access token",
			"error": err.Error(),
		})
		return
	}

	lib.SetCookies(c, "access_token", access_token, 60*30)


	c.JSON(200, gin.H{
		"message": "user signed in successfully",
		"user": gin.H{
			"id":    existingUser.ID,
			"email": existingUser.Email,
		},
	})


	// if existingUser.Password != au
	
}