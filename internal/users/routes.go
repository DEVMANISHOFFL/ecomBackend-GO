package users

import (
	"database/sql"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	// Public routes
	api.HandleFunc("/users", CreateUserController(db)).Methods("POST")
	api.HandleFunc("/login", LoginUser(db)).Methods("POST")

	// Protected routes (all authenticated users)
	protected := api.NewRoute().Subrouter()
	protected.Use(AuthMiddleware)

	// Profile – any logged-in user can view their own profile
	protected.HandleFunc("/profile/{id}", GetProfileHandler(db)).Methods("GET") //its getting access of another user from another users token

	// Role-based routes
	// Admin-only routes
	adminOnly := protected.NewRoute().Subrouter() //done ----forbidden for customers
	adminOnly.Use(RoleMiddleware("admin"))
	adminOnly.HandleFunc("/users", GetUsersController(db)).Methods("GET")
	adminOnly.HandleFunc("/users/{id}", DeleteUserController(db)).Methods("DELETE")
	adminOnly.HandleFunc("/users/{id}", UpdateUserController(db)).Methods("PUT")

	// Public: view user by id (no auth required)
	api.HandleFunc("/users/{id}", GetUserByIdController(db)).Methods("GET")
}
