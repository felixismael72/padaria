package response

import (
	"time"

	"padaria/src/core/domain"
)

type Product struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	Price          float32   `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func NewProduct(product domain.Product) *Product {
	return &Product{
		ID:             product.ID(),
		Name:           product.Name(),
		Code:           product.Code(),
		Price:          product.Price(),
		ExpirationDate: product.ExpirationDate(),
	}
}
