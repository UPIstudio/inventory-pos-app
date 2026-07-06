package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/luthfi/inventory-pos-app/backend/config"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Database berhasil konek")
			break
		}
		log.Printf("Menunggu database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Gagal konek ke database setelah 10 kali percobaan", err)
	}

}
