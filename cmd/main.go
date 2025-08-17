package main

import (
	"ecom/cmd/db"
	"log"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	log.Println("App started with DB connection")

	db.InitTables(database)
}
