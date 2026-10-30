package users

import (
	user_services "devora/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)
	
type response struct {
	ID       int64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetMe(userServices *user_services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		id, ok := value.(int64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		user, err := userServices.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve user",
			})
			return
		}

		c.JSON(http.StatusOK, response{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}
}