package image

import (
	"fmt"
	"log"
	"time"
	"github.com/product-management/db"
	"github.com/product-management/logger"
	"github.com/product-management/models"
	"github.com/streadway/amqp"
)

// ProcessProductImages simulates image processing for a product
func ProcessProductImages(productID int, productImages []string) {
	// Simulate image download and compression
	log.Printf("Processing images for product %d", productID)
	time.Sleep(2 * time.Second) // Simulating processing time

	// For simplicity, we'll pretend that the processed images are stored in S3
	processedImages := []string{"s3://bucket/processed_image1.jpg", "s3://bucket/processed_image2.jpg"}

	// Save the compressed images in the database
	err := db.UpdateCompressedImages(productID, processedImages)
	if err != nil {
		logger.Error("Error updating compressed images", err)
	}
}

// RabbitMQ Consumer to process image processing tasks
func ProcessQueueMessages() {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		// Process the image (in a real-world case, this would be the task)
		// Assume message body contains product ID and image URLs
	}
}
