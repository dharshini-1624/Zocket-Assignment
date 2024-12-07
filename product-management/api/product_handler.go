package api

import (
	"encoding/json"
	"net/http"
	"github.com/product-management/db"
	"github.com/product-management/cache"
	"github.com/product-management/image"
	"github.com/product-management/logger"
	"github.com/gorilla/mux"
	"github.com/product-management/models"
)

// CreateProduct handles the creation of a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	// Decode the request body into the product model
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logger.Error("Error decoding request body", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save product to DB
	productID, err := db.SaveProduct(product)
	if err != nil {
		logger.Error("Error saving product to DB", err)
		http.Error(w, "Failed to save product", http.StatusInternalServerError)
		return
	}

	// Push image processing task to the queue
	go image.ProcessProductImages(productID, product.ProductImages)

	// Respond with created product details
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"product_id": productID})
}

// GetProductByID retrieves a product by ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Check cache first
	product, found := cache.GetProductFromCache(id)
	if found {
		logger.Info("Cache hit for product ID", id)
	} else {
		// If not in cache, fetch from DB
		product, err := db.GetProductByID(id)
		if err != nil {
			logger.Error("Error fetching product from DB", err)
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		// Store the product in the cache for future use
		cache.SetProductToCache(id, product)
	}

	// Respond with the product details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProducts retrieves all products with optional filtering by name and price
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	name := r.URL.Query().Get("name")
	priceMin := r.URL.Query().Get("price_min")
	priceMax := r.URL.Query().Get("price_max")

	// Fetch products from DB with optional filtering
	products, err := db.GetProducts(name, priceMin, priceMax)
	if err != nil {
		logger.Error("Error fetching products", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	// Respond with the list of products
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
