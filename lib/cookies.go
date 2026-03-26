package lib

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetCookies(c *gin.Context, name, value string, maxAge int) {
	// The cookie should be secure (HTTPS only) if we are NOT in development
	isSecure := os.Getenv("GO_ENV") != "development"

	c.SetCookie(
		name,
		value,
		maxAge,
		"/", // path
		"",  // domain
		isSecure,
		true, // httpOnly
	)
}
