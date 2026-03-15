package main

import (
	"fmt"
	"net/http"

	"car-rental-api/database"
	"car-rental-api/routes"
)

func main() {

	database.ConnectDB()

	r := routes.SetupRoutes()

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", r)
}
