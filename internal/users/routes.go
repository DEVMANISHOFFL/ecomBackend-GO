package users

import (
	"database/sql"
	"ecom/internal/auth"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", CreateUserController(db)).Methods("POST")
	api.HandleFunc("/login", LoginUser(db)).Methods("POST")

	protected := api.NewRoute().Subrouter()
	protected.Use(auth.AuthMiddleware)

	protected.HandleFunc("/profile/{id}", GetProfileController(db)).Methods("GET")

	adminOnly := protected.NewRoute().Subrouter()
	adminOnly.Use(auth.RoleMiddleware("admin"))
	adminOnly.HandleFunc("/users", GetUsersController(db)).Methods("GET")
	adminOnly.HandleFunc("/users/{id}", DeleteUserController(db)).Methods("DELETE")
	adminOnly.HandleFunc("/users/{id}", UpdateUserController(db)).Methods("PUT")

	api.HandleFunc("/users/{id}", GetUserByIdController(db)).Methods("GET")
}
