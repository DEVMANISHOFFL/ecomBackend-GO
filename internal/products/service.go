package products

import (
	"database/sql"
	"errors"
)

func CreateProductService(db *sql.DB, u Product) (Product, error) {
	if u.Name == "" || u.Price == 0 {
		return Product{}, errors.New("name and price are required")
	}
	if u.Category == "" {
		u.Category = string(Other)
	}

	id, err := InsertProduct(db, u)
	if err != nil {
		return Product{}, err
	}
	return Product{
		ID:          id,
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
		Quantity:    u.Quantity,
		Category:    u.Category,
		Status:      u.Status,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}, nil
}

func GetAllProductsService(db *sql.DB) ([]Product, error) {
	return FetchProducts(db)
}
