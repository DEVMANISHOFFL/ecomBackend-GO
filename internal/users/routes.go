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
	// router.HandleFunc("/users/{id}", GetUser(db)).Methods("GET")
	// router.HandleFunc("/users/{id}", DeleteUser(db)).Methods("DELETE")
	// router.HandleFunc("/users/{id}", UpdateUser(db)).Methods("PUT")
}
