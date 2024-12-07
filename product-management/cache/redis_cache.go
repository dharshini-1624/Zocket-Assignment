package cache

import (
	"github.com/go-redis/redis/v8"
	"context"
	"encoding/json"
	"os"
	"log"
	"github.com/product-management/models"
)

var client *redis.Client
var ctx = context.Background()

// InitRedis initializes the Redis client
func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

// SetProductToCache caches the product in Redis
func SetProductToCache(id string, product models.Product) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Printf("Error marshalling product to JSON: %v", err)
		return
	}

	err = client.Set(ctx, id, productJSON, 0).Err()
	if err != nil {
		log.Printf("Error setting product to Redis: %v", err)
	}
}

// GetProductFromCache retrieves a product from Redis
func GetProductFromCache(id string) (models.Product, bool) {
	val, err := client.Get(ctx, id).Result()
	if err == redis.Nil {
		return models.Product{}, false
	} else if err != nil {
		log.Printf("Error getting product from Redis: %v", err)
		return models.Product{}, false
	}

	var product models.Product
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		log.Printf("Error unmarshalling product from JSON: %v", err)
		return models.Product{}, false
	}

	return product, true
}
