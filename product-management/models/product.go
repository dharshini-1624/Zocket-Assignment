package models

// Product represents the product data structure
type Product struct {
	ProductID          int      `json:"product_id"`
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	ProductPrice       float64  `json:"product_price"`
	CompressedImages   []string `json:"compressed_product_images"`
}
