package image

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
	"github.com/product-management/db"
	"github.com/product-management/logger"
	"github.com/product-management/models"
)

// ProcessQueueMessages consumes tasks from the image processing queue
func ProcessQueueMessages() {
	// Set up RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"image-processing", // queue name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		// Process each image task (for simplicity, assume message contains product ID and image URLs)
		productID := string(msg.Body) // In reality, parse the body as needed (e.g., JSON)
		log.Printf("Processing images for product %s", productID)
		err := processImages(productID)
		if err != nil {
			logger.Error("Image processing failed for product", err)
			// Optionally, retry or send to a dead-letter queue
		}
	}
}

// processImages simulates downloading and compressing images
func processImages(productID string) error {
	// Simulate image processing logic (e.g., download and compress)
	time.Sleep(2 * time.Second)

	// Assume successful image processing
	processedImages := []string{"s3://bucket/processed_image1.jpg", "s3://bucket/processed_image2.jpg"}

	// Update the product's compressed images in the database
	err := db.UpdateCompressedImages(productID, processedImages)
	if err != nil {
		return fmt.Errorf("failed to update compressed images for product %s: %v", productID, err)
	}

	log.Printf("Successfully processed images for product %s", productID)
	return nil
}
