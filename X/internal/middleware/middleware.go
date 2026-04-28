package middleware

import (
	"errors"
	jwtutil "go-tweet/sql/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			c.Abort()
			return
		}
		token, err := jwtutil.ValidateToken(header, secretkey, true)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		claims, ok := token.Claims.(jwtlib.MapClaims)
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token claims"))
			return
		}
		userID, _ := claims["id"].(float64)
		username, _ := claims["username"].(string)
		c.Set("user_id", int64(userID))
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshTokenMiddleware(secretkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			c.Abort()
			return
		}
		token, err := jwtutil.ValidateToken(header, secretkey, false)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		claims, ok := token.Claims.(jwtlib.MapClaims)
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token claims"))
			return
		}
		userID, _ := claims["id"].(float64)
		username, _ := claims["username"].(string)
		c.Set("user_id", int64(userID))
		c.Set("username", username)
		c.Next()
	}
}