package cart

import (
	"database/sql"
	"fmt"
)

func CreateCart(db *sql.DB, c Cart) (string, error) {
	var id string
	err := db.QueryRow(`
	INSERT INTO cart (user_id,product_id,quantity,created_at,updated_at)
	VALUES ($1,$2,$3,NOW(),NOW())
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET quantity = cart.quantity + EXCLUDED.quantity
	RETURNING id`, c.UserID, c.ProductID, c.Quantity,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to insert cart item:%w", err)
	}
	return id, nil
}

func FetchCartByUserID(db *sql.DB, userID string) ([]CartResponse, error) {
	rows, err := db.Query(`
		SELECT id, user_id, product_id, quantity, created_at, updated_at
		FROM cart WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []CartResponse
	for rows.Next() {
		var c CartResponse
		if err := rows.Scan(&c.ID, &c.UserID, &c.ProductID, &c.Quantity, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		carts = append(carts, c)
	}

	return carts, nil
}
