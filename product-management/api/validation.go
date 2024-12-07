package api

import (
	"errors"
	"strings"
	"unicode"
	"github.com/product-management/models"
)

// ValidateProduct validates the fields of a product
func ValidateProduct(product models.Product) error {
	// Validate product name
	if strings.TrimSpace(product.ProductName) == "" {
		return errors.New("product name is required")
	}
	if len(product.ProductName) > 255 {
		return errors.New("product name is too long")
	}

	// Validate product description
	if strings.TrimSpace(product.ProductDescription) == "" {
		return errors.New("product description is required")
	}
	if len(product.ProductDescription) > 1000 {
		return errors.New("product description is too long")
	}

	// Validate product images
	if len(product.ProductImages) == 0 {
		return errors.New("at least one product image is required")
	}
	for _, img := range product.ProductImages {
		if !isValidURL(img) {
			return errors.New("invalid image URL format")
		}
	}

	// Validate product price
	if product.ProductPrice <= 0 {
		return errors.New("product price must be greater than zero")
	}

	return nil
}

// isValidURL checks if a string is a valid URL (basic check for now)
func isValidURL(url string) bool {
	for _, r := range url {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == ':' || r == '/' || r == '-' || r == '?') {
			return false
		}
	}
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
