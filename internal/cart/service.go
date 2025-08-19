package cart

import (
	"database/sql"
	"errors"
)

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

func GetcartByIdService(db *sql.DB, id string) (*CartResponse, error) {
	return FetchCartById(db, id)
}
