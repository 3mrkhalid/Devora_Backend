package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	auth_services "devora/internal/services/auth"
	"os"
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


		isProduction := os.Getenv("APP_ENV") == "production"
        
        c.SetCookie(
            "refresh_token",
            user.RefreshToken,
            7*24*60*60,
            "/",
            "",
            isProduction,
            true,
        )

        c.JSON(http.StatusOK, gin.H{
            "message": "Registration successful",
            "user": gin.H{
                "id": user.ID,
                "email": user.Email,
                "username": user.Username,
            },
            "accessToken": user.AccessToken,
        })

	}
}

