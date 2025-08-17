package users

import (
	"database/sql"
	"ecom/pkg/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {

	router.Use(middlewares.JSONContentTypeMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", GetUsersController(db)).Methods("GET")
	api.HandleFunc("/users", CreateUserController(db)).Methods("POST")
	api.HandleFunc("/users/{id}", GetUserByIdController(db)).Methods("GET")
	api.HandleFunc("/users/{id}", DeleteUserController(db)).Methods("DELETE")
	api.HandleFunc("/users/{id}", UpdateUserController(db)).Methods("PUT")
}
