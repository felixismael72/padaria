package primary

import "padaria/src/core/domain"

type ProductManager interface {
	RegisterProduct(product domain.Product) (int, error)
	ListProducts() ([]domain.Product, error)
	EditProduct(product domain.Product) error
	RemoveProduct(productID int) error
}
