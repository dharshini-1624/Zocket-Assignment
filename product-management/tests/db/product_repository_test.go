package db

import (
	"testing"
	"github.com/product-management/models"
	"github.com/stretchr/testify/assert"
)

// TestSaveProduct tests the SaveProduct function
func TestSaveProduct(t *testing.T) {
	// Setup a mock product
	product := models.Product{
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Call SaveProduct function
	productID, err := SaveProduct(product)
	assert.Nil(t, err)
	assert.Greater(t, productID, 0) // Ensure that productID is greater than 0
}

// TestGetProductByID tests the GetProductByID function
func TestGetProductByID(t *testing.T) {
	// Setup a mock product
	product := models.Product{
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Save the product to the DB
	productID, err := SaveProduct(product)
	assert.Nil(t, err)

	// Call GetProductByID function
	retrievedProduct, err := GetProductByID(string(productID))
	assert.Nil(t, err)
	assert.Equal(t, product.ProductName, retrievedProduct.ProductName)
}
