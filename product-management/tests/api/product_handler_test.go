package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/product-management/db"
	"github.com/product-management/models"
	"github.com/stretchr/testify/assert"
)

// TestCreateProduct tests the CreateProduct handler
func TestCreateProduct(t *testing.T) {
	// Setup a mock product
	product := models.Product{
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Marshal product into JSON
	productJSON, err := json.Marshal(product)
	assert.Nil(t, err)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	assert.Nil(t, err)

	// Setup the router and handler
	r := mux.NewRouter()
	r.HandleFunc("/products", CreateProduct).Methods("POST")

	// Record the response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert the response status code is 201 (Created)
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Assert the response body contains the product ID
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Contains(t, response, "product_id")
}

// TestGetProductByID tests the GetProductByID handler
func TestGetProductByID(t *testing.T) {
	// Setup a mock product
	product := models.Product{
		ProductID:          1,
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Mock database call (use a mock or real DB for integration tests)
	db.SaveProduct(product)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products/1", nil)
	assert.Nil(t, err)

	// Setup the router and handler
	r := mux.NewRouter()
	r.HandleFunc("/products/{id:[0-9]+}", GetProductByID).Methods("GET")

	// Record the response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert the response status code is 200 (OK)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body contains the product data
	var response models.Product
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, product.ProductID, response.ProductID)
	assert.Equal(t, product.ProductName, response.ProductName)
}

// TestGetProducts tests the GetProducts handler with query params
func TestGetProducts(t *testing.T) {
	// Setup a mock product
	product := models.Product{
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Mock database call (use a mock or real DB for integration tests)
	db.SaveProduct(product)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products?name=Test", nil)
	assert.Nil(t, err)

	// Setup the router and handler
	r := mux.NewRouter()
	r.HandleFunc("/products", GetProducts).Methods("GET")

	// Record the response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert the response status code is 200 (OK)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body contains the product data
	var products []models.Product
	err = json.NewDecoder(rr.Body).Decode(&products)
	assert.Nil(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, product.ProductName, products[0].ProductName)
}
