package dto

import (
	"time"

	"padaria/src/core/domain"
)

type Product struct {
	ID             int       `db:"product_id"`
	Name           string    `db:"product_name"`
	Code           string    `db:"product_code"`
	Price          float32   `db:"product_price"`
	ExpirationDate time.Time `db:"product_expiration_date"`
}

func (dto Product) ToDomain() *domain.Product {
	return domain.NewProduct(dto.ID, dto.Name, dto.Code, dto.Price, dto.ExpirationDate)
}
