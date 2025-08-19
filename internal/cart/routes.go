package cart

import (
	"database/sql"
	"ecom/internal/auth"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	protected := api.NewRoute().Subrouter()
	protected.Use(auth.AuthMiddleware)
	adminOnly := protected.NewRoute().Subrouter()
	adminOnly.Use(auth.RoleMiddleware("admin"))

	protected.HandleFunc("/cart/{id}", GetCartByIdController(db)).Methods("GET")
	protected.HandleFunc("/cart", CreateCartController(db)).Methods("POST")
	// api.HandleFunc("/cart/{id}", DeleteCartController(db)).Methods("DELETE")
	// api.HandleFunc("/cart/{id}", UpdateCartController(db)).Methods("PUT")

}
