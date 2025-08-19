package db

import (
	"database/sql"
	"log"
)

func InitCartTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS cart (
	    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	    quantity INT NOT NULL DEFAULT 1 CHECK (quantity > 0),
	    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	    UNIQUE(user_id, product_id)
	)`)
	if err != nil {
		log.Fatal("Error creating cart table:", err)
	}
}
