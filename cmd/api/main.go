package main

import (
	"log"
	"net/http"
	"os"

	"devora/internal/platform/database"
	"devora/internal/middlewares"
	repositories "devora/internal/repositories/user"
	"devora/internal/services"

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

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// TODO: Use userService in your handlers.

	_ = userService

	router.GET("/", welcome)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}