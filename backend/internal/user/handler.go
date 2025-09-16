package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserHandler handles user registration
func CreateUserHandler(repo *UserRepository, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Temporary struct to receive password from request
	var req struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to hash password"})
		return
	}

	// Map request to User struct
	now := time.Now()
	u := User{
		Email:               req.Email,
		Username:            req.Username,
		FirstName:           req.FirstName,
		LastName:            req.LastName,
		PasswordHash:        string(hashedPassword),
		IsActive:            true,
		IsVerified:          false,
		CreatedAt:           now,
		UpdatedAt:           now,
		FailedLoginAttempts: 0,
	}

	// Save user using repository
	if err := repo.CreateUser(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		// Handle duplicate entry error for MySQL
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			json.NewEncoder(w).Encode(map[string]string{"error": "Email or username already exists"})
			return
		}

		// Return the actual error message
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginHandler handles user login
func LoginHandler(repo *UserRepository, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Find user by email
	u, err := repo.GetUserByEmail(creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email or password"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email or password"})
		return
	}

	// Update last login timestamp
	now := time.Now()
	u.LastLogin = &now
	_ = repo.UpdateUser(u)

	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

