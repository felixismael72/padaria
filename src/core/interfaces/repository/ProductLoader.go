package repository

import "padaria/src/core/domain"

type ProductLoader interface {
	InsertProduct(product domain.Product) (int, error)
}
