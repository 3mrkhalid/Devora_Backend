package auth

import (
	auth_services "devora/internal/services/auth"
	"net/http"

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

        // Return a safe user response; do not expose the password hash.
        _ = user

        //Access token and Refresh token LoginHandler
        // return a respone {"message":{{"login success", Access_Token}}
    }
}
