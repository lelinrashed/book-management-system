package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lelinrashed/go-bookstore/pkg/routes"
)

func main() {
	// Initialize route
	r := mux.NewRouter()

	// Add book store route
	routes.RegisterBookStoreRoutes(r)

	// Handle home route
	http.Handle("/", r)

	// Run the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
