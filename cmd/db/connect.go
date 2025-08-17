package db

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	return db
}
