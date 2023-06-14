package repository

import "padaria/src/core/domain"

type ProductLoader interface {
	InsertProduct(product domain.Product) (int, error)
	SelectProducts() ([]domain.Product, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(productID int) error
}
