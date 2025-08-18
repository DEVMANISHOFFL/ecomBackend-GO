package cart

import (
	"database/sql"
	"fmt"
)

func CreateCart(db *sql.DB, c Cart) (string, error) {
	var id string
	err := db.QueryRow(`
	INSERT INTO cart_items (user_id,product_id,quantity,created_at,updated_at)
	VALUES ($1,$2,$3,NOW(),NOW())
	RETURNING id`, c.UserID, c.ProductID, c.Quantity,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to insert cart item:%w", err)
	}
	return id, nil
}
