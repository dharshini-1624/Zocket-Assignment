package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/product-management/api"
	"github.com/product-management/db"
	"github.com/product-management/cache"
	"github.com/product-management/logger"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize components
	logger.Init()
	db.InitDB()
	cache.InitRedis()

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/products", api.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", api.GetProductByID).Methods("GET")
	r.HandleFunc("/products", api.GetProducts).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
