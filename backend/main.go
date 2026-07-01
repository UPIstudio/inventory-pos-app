package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	host := "db"
	port := 5432
	user := "user"
	password := "password"
	dbname := "inventory_db"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
