package products

import (
	"database/sql"
	"fmt"
)

func InsertProduct(db *sql.DB, u Product) (string, error) {
	var id string
	err := db.QueryRow(`
	INSERT INTO products (name, description,price,quantity,category,status,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,NOW(),NOW())
	RETURNING id`, u.Name, u.Description, u.Price, u.Quantity, u.Category, u.Status).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to insert product: %w", err)
	}
	return id, err
}

func FetchProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id,name,description,price,quantity,category,status,created_at,updated_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var u Product
		if err := rows.Scan(&u.ID, &u.Name, &u.Description, &u.Price, &u.Quantity, &u.Category, &u.Status, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, err
}

func FetchProductById(db *sql.DB, id string) (*Product, error) {
	var u Product
	err := db.QueryRow("SELECT id,name,description,price,quantity,category,status,created_at,updated_at FROM products WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Description, &u.Price, &u.Quantity, &u.Category, &u.Status, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &u, nil
}

func DeleteProduct(db *sql.DB, id string) (bool, error) {
	res, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}
	rowsaffected, _ := res.RowsAffected()
	if rowsaffected == 0 {
		return false, nil
	}
	return true, nil
}

func UpdateProduct(db *sql.DB, id string, u Product) (*Product, error) {
	_, err := db.Exec(`
	UPDATE products
	SET name=$1, description=$2, price=$3, quantity=$4, category=$5, status=$6, updated_at=NOW()
	WHERE id=$7`, u.Name, u.Description, u.Price, u.Quantity, u.Category, u.Status, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	return FetchProductById(db, id)
}
