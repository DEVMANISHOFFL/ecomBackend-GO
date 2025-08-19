package cart

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserCart struct {
	UserID string     `json:"user_id"`
	Items  []CartItem `json:"items"`
}

func CreateCartService(db *sql.DB, c Cart) (CartResponse, error) {
	if c.UserID == "" || c.ProductID == "" {
		return CartResponse{}, errors.New("user_id and product_id are required")
	}
	if c.Quantity <= 0 {
		return CartResponse{}, errors.New("quantity must be atleast 1")
	}
	id, err := CreateCart(db, c)
	if err != nil {
		return CartResponse{}, err
	}
	return CartResponse{
		ID:        id,
		UserID:    c.UserID,
		ProductID: c.ProductID,
		Quantity:  c.Quantity,
	}, nil
}

func GetCartByUserIDService(db *sql.DB, userID string) (*UserCart, error) {
	rows, err := db.Query(`
		SELECT id, product_id, quantity, created_at, updated_at
		FROM cart
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %w", err)
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &UserCart{
		UserID: userID,
		Items:  items,
	}, nil
}
