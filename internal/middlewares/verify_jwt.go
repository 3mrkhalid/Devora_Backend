
package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenParts := strings.Fields(authHeader)

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "missing or invalid token"},
			)
			return
		}

		secretKey := os.Getenv("AccessToken")
		if secretKey == "" {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": "authentication configuration error"},
			)
			return
		}

		claims := jwt.MapClaims{}

		user, err := jwt.ParseWithClaims(
			tokenParts[1],
			&claims,
			func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, errors.New("unexpected signing method")
				}

				return []byte(secretKey), nil
			},
		)

		if err != nil || user == nil || !user.Valid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid or expired token"},
			)
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
