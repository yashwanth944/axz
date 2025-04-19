package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/user/apigateway/auth"
	"github.com/user/apigateway/db"
	"github.com/user/apigateway/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Set up database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Debug endpoint for password testing
	router.POST("/debug/check-password", func(c *gin.Context) {
		var data struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUserByEmail(database, data.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found",
				"email": data.Email,
			})
			return
		}

		// Manually compare passwords
		result := auth.CheckPasswordHash(data.Password, user.Password)

		c.JSON(http.StatusOK, gin.H{
			"passwordMatches": result,
			"storedHash":      user.Password,
			"inputPassword":   data.Password,
		})
	})

	// After debug/check-password endpoint, add this new route
	router.POST("/debug/create-test-user", func(c *gin.Context) {
		// Create a test user with a known password
		testUser := models.User{
			Email:     "test123@example.com",
			Password:  "password123",
			FirstName: "Test",
			LastName:  "User",
		}

		if err := db.CreateUser(database, &testUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create test user",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Test user created successfully",
			"user":    testUser.SafeResponse(),
			"login_credentials": map[string]string{
				"email":    testUser.Email,
				"password": "password123",
			},
		})
	})

	// Add after the other debug endpoints
	router.POST("/debug/verify-token", func(c *gin.Context) {
		var tokenData struct {
			Token string `json:"token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&tokenData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse the token without validation
		token, _ := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
			return []byte("dummy"), nil
		})

		// Get the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.JSON(http.StatusOK, gin.H{
				"claims": claims,
				"raw":    tokenData.Token,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse token claims"})
		}
	})

	// Public routes
	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.CreateUser(database, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUserByEmail(database, loginData.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if !auth.CheckPasswordHash(loginData.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token
		token, err := auth.GenerateJWT(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	})

	// Protected routes group
	protected := router.Group("/api")
	protected.Use(auth.JWTMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			// Get user ID from context (set by middleware)
			userID, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			user, err := db.GetUserByID(database, userID.(uint))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user.SafeResponse()})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Server running on port %s\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
