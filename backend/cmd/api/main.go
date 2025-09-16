package main

import (
    "net/http"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

// Example User struct
type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// In-memory users slice for demo
var users = []User{}

func main() {
    r := gin.Default()

    // Enable CORS for all origins (development only)
    r.Use(cors.Default())

    // Routes
    r.POST("/register", registerHandler)
    r.POST("/login", loginHandler)

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "API is running"})
    })

    r.Run(":8080") // listen on port 8080
}

// Register handler
func registerHandler(c *gin.Context) {
    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users = append(users, u)
    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": u})
}

// Login handler
func loginHandler(c *gin.Context) {
    var loginData User
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Simple login check
    for _, u := range users {
        if u.Email == loginData.Email && u.Password == loginData.Password {
            c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": u})
            return
        }
    }
    c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
}

