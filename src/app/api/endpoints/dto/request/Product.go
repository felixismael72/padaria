package request

import (
	"padaria/src/core/domain"
	"time"
)

type Product struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func (dto Product) ToDomain() *domain.Product {
	return domain.NewProduct(0, dto.Name, dto.Code, dto.Price, time.Time{})
}
