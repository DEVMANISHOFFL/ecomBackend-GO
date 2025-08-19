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

func FetchCartById(db *sql.DB, id string) (*CartResponse, error) {
	var u CartResponse
	err := db.QueryRow("SELECT id,user_id,product_id,quantity,created_at,updated_at FROM cart WHERE id = $1", id).Scan(&u.ID, &u.UserID, &u.ProductID, &u.Quantity, &u.CreatedAt, &u.UpdatedAt)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &u, nil
}
