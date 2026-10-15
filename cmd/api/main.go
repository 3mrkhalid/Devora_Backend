
package main

import (
	"log"
	"net/http"
	"os"

	"devora/internal/middlewares"
	"devora/internal/platform/database"
	auth_repositories "devora/internal/repositories/auth"
	user_repositories "devora/internal/repositories/user"
	auth_services "devora/internal/services/auth"
	user_services "devora/internal/services/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Devora",
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system environment")
	}

	db, err := database.NewMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(middlewares.NewCORS())

	userRepo := user_repositories.NewUserRepository(db)
	userService := user_services.NewUserService(userRepo)

	loginRepo := auth_repositories.NewLoginRepository(db)
	authService := auth_services.NewAuthService(loginRepo)

	// TODO: Pass userService and authService to their handlers.
	_, _ = userService, authService

	router.GET("/", welcome)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
