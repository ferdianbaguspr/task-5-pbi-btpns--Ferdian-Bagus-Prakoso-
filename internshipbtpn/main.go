package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" gorm:"not null;unique"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Photos    []Photo
}

type Photo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

func main() {
	// Membuat koneksi ke database SQLite
	db, err = gorm.Open("sqlite3", "your_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate database
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Photo{})

	r := gin.Default()

	// Endpoint untuk registrasi user
	r.POST("/users/register", createUser)

	// Endpoint untuk login user
	r.GET("/users/login", loginUser)

	// Endpoint untuk mengupdate user
	r.PUT("/users/:userId", updateUser)

	// Endpoint untuk menghapus user
	r.DELETE("/users/:userId", deleteUser)

	// Endpoint untuk membuat foto
	r.POST("/photos", createPhoto)

	// Endpoint untuk mendapatkan semua foto
	r.GET("/photos", getPhotos)

	// Endpoint untuk mengupdate foto
	r.PUT("/photos/:photoId", updatePhoto)

	// Endpoint untuk menghapus foto
	r.DELETE("/photos/:photoId", deletePhoto)

	// Menjalankan server
	r.Run(":8080")
}
