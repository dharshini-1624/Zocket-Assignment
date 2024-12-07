package cache

import (
	"testing"
	"github.com/product-management/models"
	"github.com/stretchr/testify/assert"
)

// TestSetProductToCache tests the caching of a product
func TestSetProductToCache(t *testing.T) {
	product := models.Product{
		ProductID:          1,
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}

	// Set the product to cache
	SetProductToCache("1", product)

	// Retrieve the product from cache
	cachedProduct, found := GetProductFromCache("1")
	assert.True(t, found)
	assert.Equal(t, product.ProductID, cachedProduct.ProductID)
	assert.Equal(t, product.ProductName, cachedProduct.ProductName)
}

// TestGetProductFromCache tests the retrieval of a product from cache
func TestGetProductFromCache(t *testing.T) {
	// Retrieve a non-existent product
	_, found := GetProductFromCache("999")
	assert.False(t, found)
}
