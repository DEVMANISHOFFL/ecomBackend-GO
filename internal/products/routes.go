package products

import (
	"database/sql"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {

	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/products", GetProductsController(db)).Methods("GET")
	api.HandleFunc("/products", CreateProductController(db)).Methods("POST")
	api.HandleFunc("/products/{id}", GetProductByIdController(db)).Methods("GET")
	api.HandleFunc("/products/{id}", DeleteProductController(db)).Methods("DELETE")
	api.HandleFunc("/products/{id}", UpdateProductController(db)).Methods("PUT")
}
