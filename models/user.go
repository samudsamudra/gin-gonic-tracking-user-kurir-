package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User - Struktur data untuk tabel users
type User struct {
	gorm.Model              // Menambahkan kolom ID, CreatedAt, UpdatedAt, DeletedAt otomatis
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"type:enum('user','admin');default:'user'" json:"role"`
}

// HashPassword - Fungsi untuk meng-hash password sebelum disimpan
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Gagal melakukan hash password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword - Fungsi untuk memeriksa password dengan hash di database
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CreateUser - Menambahkan user baru ke database
func CreateUser(db *gorm.DB, user *User) error {
	// Hash password sebelum disimpan
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Simpan ke database
	if err := db.Create(user).Error; err != nil {
		log.Println("Gagal menambahkan user ke database:", err)
		return err
	}
	return nil
}

// GetUserByEmail - Mengambil user berdasarkan email
func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User tidak ditemukan:", err)
		} else {
			log.Println("Gagal mengambil user dari database:", err)
		}
		return nil, err
	}
	return &user, nil
}
