
package routes

import (
	auth "devora/internal/controllers/auth"
	auth_services "devora/internal/services/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(
	router *gin.RouterGroup,
	authService *auth_services.AuthService,
) {
	authGroup := router.Group("/auth")

	{
		authGroup.POST("/register", auth.Register(authService))
		authGroup.POST("/login", auth.LoginHandler(authService))
		authGroup.POST("/logout", auth.Logout)
	}
}
