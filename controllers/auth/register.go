package controllers

import (
	"database/sql"
	"dealsafe/database"
	dealsafe "dealsafe/database/sqlc"
	"dealsafe/lib"
	request "dealsafe/types"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



func RegisterUser(c *gin.Context) {
	var authRequest request.AuthRequest

	if err := c.BindJSON(&authRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	existingUser, err := database.Queries.FindUserByEmail(
		c,
		authRequest.Email,
	)
	
	// Check if there was a database error other than "user not found"
	if err != nil && err != sql.ErrNoRows {
		c.JSON(500, gin.H{"message": "Error checking existing user", "error": err.Error()})
		return
	}

	if existingUser.ID != uuid.Nil {
		c.JSON(400, gin.H{
			"message": "User already exists",
		})
		return
	}

	pubKey, privKey, err := lib.GenerateAPIKeys()
	if err != nil {
		c.JSON(500, gin.H{"message": "Error generating API keys", "error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := database.Queries.CreateUser(
		c,
		dealsafe.CreateUserParams{
			Email:    authRequest.Email,
			Password: string(hashedPassword),
			TestPubKey: sql.NullString{
				String: pubKey,
				Valid:  true,
			},
			TestPrivKey: sql.NullString{
				String: privKey,
				Valid:  true,
			},
			TermsAccepted: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		},
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error": err.Error(),
		})
		return
	}


	access_token, err := lib.GenerateJWT(user.Email, user.ID.String())

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to generate access token",
			"error": err.Error(),
		})
		return
	}

	lib.SetCookies(c, "access_token", access_token, 60*30)

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user": user,
	})
	
}