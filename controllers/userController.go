package controllers

import (
	"net/http"
	"time"

	"sistem-tracking/config"
	"sistem-tracking/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key untuk JWT
var jwtSecret = []byte("secret-key")

// RegisterUser - Endpoint untuk mendaftarkan user baru
func RegisterUser(c *gin.Context) {
	// Struct untuk validasi input
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role" binding:"required,oneof=user admin"`
	}

	// Bind input JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
		return
	}

	// Buat user baru dengan hash password
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	if err := models.CreateUser(config.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan user ke database"})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{
		"message": "User berhasil didaftarkan",
		"email":   input.Email,
		"role":    input.Role,
	})
}

// LoginUser - Endpoint untuk autentikasi pengguna
func LoginUser(c *gin.Context) {
	// Struct untuk validasi input
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind input JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
		return
	}

	// Ambil user berdasarkan email
	user, err := models.GetUserByEmail(config.DB, input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	// Verifikasi password
	if err := models.CheckPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	// Generate token JWT
	token, err := generateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}

// generateJWT - Fungsi untuk membuat token JWT
func generateJWT(email, role string) (string, error) {
	// Buat claims untuk token
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}

	// Buat token menggunakan HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
