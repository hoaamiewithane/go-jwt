package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/hoaamiewithane/go-jwt/routes"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go JWT App",
		})
	})
    err = router.Run(":" + port)
    if err != nil {
        return
    }
}
