package db

import (
	"database/sql"
	"log"
)

func InitTables(db *sql.DB) {
	_, err := db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS users (
	    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	    name TEXT NOT NULL,
	    email TEXT UNIQUE NOT NULL,
	    password TEXT NOT NULL,
	    role VARCHAR(20) NOT NULL DEFAULT 'customer' CHECK (role IN ('customer','admin','seller')),
	    created_at TIMESTAMP DEFAULT NOW(),
	    updated_at TIMESTAMP DEFAULT NOW()
	)`)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}
}
