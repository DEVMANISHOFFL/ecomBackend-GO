package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	return db
}
