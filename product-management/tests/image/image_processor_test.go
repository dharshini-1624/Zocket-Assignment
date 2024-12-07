package image

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/product-management/db"
	"github.com/product-management/models"
)

// TestProcessImages tests the image processing logic
func TestProcessImages(t *testing.T) {
	// Setup a mock product ID for testing
	productID := "1"
	
	// Simulate saving a product
	product := models.Product{
		ProductID:          1,
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "A sample test product",
		ProductImages:      []string{"http://example.com/image1.jpg"},
		ProductPrice:       19.99,
	}
	err := db.SaveProduct(product)
	assert.Nil(t, err)

	// Process the images (simulated)
	err = processImages(productID)
	assert.Nil(t, err)

	// Verify the images were processed and updated in the DB
	processedImages, err := db.GetProcessedImages(productID)
	assert.Nil(t, err)
	assert.NotEmpty(t, processedImages)
	assert.Len(t, processedImages, 2) // Simulating 2 processed images
}
