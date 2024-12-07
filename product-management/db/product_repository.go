package db

import (
	"fmt"
	"github.com/product-management/models"
)

// SaveProduct saves a new product to the database
func SaveProduct(product models.Product) (int, error) {
	var id int
	err := db.QueryRow(
		`INSERT INTO products (user_id, product_name, product_description, product_images, product_price) 
		VALUES ($1, $2, $3, $4, $5) RETURNING product_id`,
		product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error saving product: %v", err)
	}
	return id, nil
}

// GetProductByID retrieves a product by ID
func GetProductByID(id string) (models.Product, error) {
	var product models.Product
	err := db.QueryRow(
		`SELECT product_id, user_id, product_name, product_description, product_images, product_price 
		FROM products WHERE product_id = $1`,
		id,
	).Scan(&product.ProductID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice)
	if err != nil {
		return product, fmt.Errorf("error retrieving product: %v", err)
	}
	return product, nil
}

// GetProducts fetches products with optional filters
func GetProducts(name, priceMin, priceMax string) ([]models.Product, error) {
	var query string
	var args []interface{}
	query = `SELECT product_id, user_id, product_name, product_description, product_images, product_price FROM products WHERE 1=1`

	if name != "" {
		query += " AND product_name ILIKE $1"
		args = append(args, "%"+name+"%")
	}
	if priceMin != "" {
		query += " AND product_price >= $2"
		args = append(args, priceMin)
	}
	if priceMax != "" {
		query += " AND product_price <= $3"
		args = append(args, priceMax)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ProductID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice); err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		products = append(products, product)
	}
	return products, nil
}
