package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing auth token"})
			return
		}

		tokenString = tokenString[7:]

		token, err := ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired auth token"})
			return
		}

		var wl WhiteList
		wl.Create()
		defer wl.Close()

		ID := strconv.FormatUint(uint64(GetTokenClaims(token)["id"].(float64)), 10)

		wlToken, err := wl.HGet(c, ID, "token")

		if wlToken != tokenString || !token.Valid || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired auth token"})
			return
		}

		c.Set("claims", GetTokenClaims(token))
		c.Set("id", GetTokenClaims(token)["id"])

		c.Next()
	}
}
