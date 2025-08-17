package db

import (
	"database/sql"
	"log"
)

func InitProductTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS products (
	    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	    name TEXT NOT NULL,
	    description TEXT,
	    price NUMERIC(10,2) NOT NULL,
	    quantity INT NOT NULL,
	    category VARCHAR(50) NOT NULL CHECK (category IN ('Electronics','Clothing','Home','Books')),
	    status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active','inactive','out_of_stock','discontinued')),
	    created_at TIMESTAMP DEFAULT NOW(),
	    updated_at TIMESTAMP DEFAULT NOW()
	)`)
	if err != nil {
		log.Fatal("Error creating products table:", err)
	}
}
