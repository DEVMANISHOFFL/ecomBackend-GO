package main

import (
	"ecom/cmd/db"
	"ecom/internal/cart"
	"ecom/internal/products"
	"ecom/internal/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	log.Println("App started with DB connection")

	db.InitUserTable(database)
	db.InitProductTable(database)
	db.InitCartTable(database)

	router := mux.NewRouter()
	users.RegisterRoutes(router, database)
	products.RegisterRoutes(router, database)
	cart.RegisterRoutes(router, database)

	log.Println("Server started at http://localhost:8082")
	if err := http.ListenAndServe(":8082", router); err != nil {
		log.Fatal(err)
	}
}
