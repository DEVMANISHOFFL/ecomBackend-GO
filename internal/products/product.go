package products

import "time"

type ProductStatus string
type Category string

const (
	StatusActive     ProductStatus = "active"
	StatusInactive   ProductStatus = "inactive"
	StatusOutOfStock ProductStatus = "out_of_stock"
	Discontinued     ProductStatus = "discontinued"
)

const (
	CategoryElectronics Category = "Electronics"
	CategoryClothing    Category = "Clothing"
	CategoryHome        Category = "Home"
	CategoryBooks       Category = "Books"
	Other               Category = "Other"
)

type Product struct {
	ID          string        `json:"id" db:"id"`
	Name        string        `json:"name" db:"name"`
	Description string        `json:"description" db:"description"`
	Price       float64       `json:"price" db:"price"`
	Quantity    int           `json:"quantity" db:"quantity"`
	Category    string        `json:"category" db:"category"`
	Status      ProductStatus `json:"status" db:"status"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}
