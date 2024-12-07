package api

import (
	"github.com/product-management/db"
	"github.com/product-management/models"
)

// CreateProductService handles the business logic for creating a product
func CreateProductService(product models.Product) (int, error) {
	return db.SaveProduct(product)
}

// GetProductService fetches a product by its ID
func GetProductService(id string) (models.Product, error) {
	return db.GetProductByID(id)
}

// GetProductsService fetches products with optional filtering
func GetProductsService(name, priceMin, priceMax string) ([]models.Product, error) {
	return db.GetProducts(name, priceMin, priceMax)
}
