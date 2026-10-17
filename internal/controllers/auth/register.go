package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	auth_services "devora/internal/services/auth"
)


type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8"`
}


func Register(authService *auth_services.AuthService) gin.HandlerFunc {
	return func (c *gin.Context)  {

		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := authService.CheckEmail(c.Request.Context(), req.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := authService.RegisterUser(c.Request.Context(), req.Username, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}


		c.JSON(http.StatusAccepted, gin.H{
			"data": gin.H{
				"message": "user registered successfully",
				"user": gin.H{
					"id": user.ID,
					"username": user.Username,
					"email": user.Email,
				},
			},
		})
	}
}

