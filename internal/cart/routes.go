package cart

import (
	"database/sql"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	// api.HandleFunc("/cart", GetCartController(db)).Methods("GET")
	api.HandleFunc("/cart", CreateCartController(db)).Methods("POST")
	// api.HandleFunc("/cart/{id}", DeleteCartController(db)).Methods("DELETE")
	// api.HandleFunc("/cart/{id}", UpdateCartController(db)).Methods("PUT")

}
