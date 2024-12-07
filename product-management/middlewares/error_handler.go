package image

import (
	"time"
	"fmt"
	"log"
)

// Retry logic for image processing
func RetryImageProcessing(productID int, retries int) {
	for i := 0; i < retries; i++ {
		err := processImage(productID)
		if err == nil {
			log.Printf("Image processed successfully for product %d", productID)
			return
		}
		log.Printf("Error processing image for product %d. Retry %d/%d", productID, i+1, retries)
		time.Sleep(time.Duration(i+1) * time.Second) // Exponential backoff
	}
	log.Printf("Failed to process image for product %d after %d retries", productID, retries)
}

func processImage(productID int) error {
	// Simulate image processing
	if productID%2 == 0 {
		return nil // success
	}
	return fmt.Errorf("random error processing product %d", productID)
}
