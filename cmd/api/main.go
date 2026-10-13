package main

import (
	"log"
	"net/http"
	"os"

	"devora/internal/middlewares"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)
	
	func getting(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
	
	func main() {

		if err := godotenv.Load(); err != nil {
			log.Fatal("No .env file found; using system environment")
		}
		
		router := gin.Default()

	    router.Use(middleware.NewCORS())

	    port := os.Getenv("PORT")
					
		router.Run(port)
	}