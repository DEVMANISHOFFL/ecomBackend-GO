package users

import (
	"database/sql"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	// Middleware to set JSON headers
	router.Use(middlewares.JSONContentTypeMiddleware)

	// API v1 prefix
	api := router.PathPrefix("/api/v1").Subrouter()

	// Protected routes (require JWT)
	protected := api.NewRoute().Subrouter()
	protected.Use(AuthMiddleware)

	// Public routes
	api.HandleFunc("/users", CreateUserController(db)).Methods("POST")
	api.HandleFunc("/login", LoginUser(db)).Methods("POST")

	// Protected routes
	protected.HandleFunc("/profile/{id}", GetProfileHandler(db)).Methods("GET")
	protected.HandleFunc("/users", GetUsersController(db)).Methods("GET")

	// Public user by id
	// Regex ensures only valid UUIDs are matched
	api.HandleFunc("/users/{id}", GetUserByIdController(db)).Methods("GET")
	api.HandleFunc("/users/{id}", DeleteUserController(db)).Methods("DELETE")
	api.HandleFunc("/users/{id}", UpdateUserController(db)).Methods("PUT")
}
