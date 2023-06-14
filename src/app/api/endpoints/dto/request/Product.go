package request

import (
	"time"

	"padaria/src/core/domain"
)

type Product struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func (dto Product) ToDomain() *domain.Product {
	return domain.NewProduct(0, dto.Name, dto.Code, dto.Price, time.Time{})
}

func (dto Product) ToDomainWithID(id int) *domain.Product {
	return domain.NewProduct(id, dto.Name, dto.Code, dto.Price, time.Time{})
}
