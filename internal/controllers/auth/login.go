package auth

import (
	auth_services "devora/internal/services/auth"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=254"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

func LoginHandler(authService *auth_services.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {

    	var req LoginRequest
     	
     	if err := c.BindJSON(&req); err != nil {
      		c.JSON(http.StatusBadRequest, gin.H {
        		"error":"Invalid Email or Password Format",
        	})
        	return
      	}
       
        user, err := authService.VerifyLogin(
            c.Request.Context(),
            req.Email,
            req.Password,
        )
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid email or password",
            })
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
            "message": "Login successful",
            "user": gin.H{
                "id": user.ID,
                "email": user.Email,
                "username": user.Username,
            },
            "accessToken": user.AccessToken,
        })
    }
}
