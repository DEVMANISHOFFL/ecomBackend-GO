package main

import (
	"ecom/cmd/db"
	"ecom/internal/cart"
	"ecom/internal/products"
	"ecom/internal/users"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	// router.Use(middlewares.CORSMiddleware)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Println("Server started at http://localhost:8082")
	if err := http.ListenAndServe(":8082", handlers.CORS(originsOk, headersOk, methodsOk)(router)); err != nil {
		log.Fatal(err)
	}
}
