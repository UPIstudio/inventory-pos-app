package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/luthfi/inventory-pos-app/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Menunggu database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Gagal konek ke databse: ", err)
	}

	db.AutoMigrate(&models.Product{}, &models.User{})

	DB = db

	fmt.Println("Database terkoneksi dan migrasi selesai!")

}
