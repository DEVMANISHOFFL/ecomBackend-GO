package products

import (
	"database/sql"
	"ecom/internal/users"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {

	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/products", GetProductsController(db)).Methods("GET")
	api.HandleFunc("/products/{id}", GetProductByIdController(db)).Methods("GET")

	protected := api.NewRoute().Subrouter()
	protected.Use(users.AuthMiddleware)
	adminOnly := protected.NewRoute().Subrouter()
	adminOnly.Use(users.RoleMiddleware("admin"))
	adminOnly.HandleFunc("/products", CreateProductController(db)).Methods("POST")
	adminOnly.HandleFunc("/products/{id}", DeleteProductController(db)).Methods("DELETE")
	adminOnly.HandleFunc("/products/{id}", UpdateProductController(db)).Methods("PUT")
}
