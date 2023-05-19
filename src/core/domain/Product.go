package domain

import "time"

type Product struct {
	id             int
	name           string
	code           string
	price          float32
	expirationDate time.Time
}

func (p Product) ID() int {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Code() string {
	return p.code
}

func (p Product) Price() float32 {
	return p.price
}

func (p Product) ExpirationDate() time.Time {
	return p.expirationDate
}

func NewProduct(id int, name, code string, price float32, expirationDate time.Time) *Product {
	return &Product{
		id:             id,
		name:           name,
		code:           code,
		price:          price,
		expirationDate: expirationDate,
	}
}
