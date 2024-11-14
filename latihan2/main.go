package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisiasi Gin router
	router := gin.Default()

	// Middleware: Logger
	router.Use(gin.Logger())

	// Middleware: Recovery
	router.Use(gin.Recovery())

	// Route definition
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/halo/:name", func(c *gin.Context) { // Perbaikan: gunakan `:name` agar Gin bisa mengambil parameter dari URL
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Halo, " + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		if loginData.Email == "example@example.com" && loginData.Password == "password123" {
			c.JSON(200, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
	})

	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Name parameter is missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	router.Run(":8080")
}
