package services

import (
	"log"
	"padaria/src/core/domain"
	"padaria/src/core/interfaces/primary"
	"padaria/src/core/interfaces/repository"
)

var _ primary.ProductManager = (*ProductServices)(nil)

type ProductServices struct {
	productRepository repository.ProductLoader
}

func (service ProductServices) RegisterProduct(product domain.Product) (int, error) {
	productID, err := service.productRepository.InsertProduct(product)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return productID, nil
}

func NewProductServices(productRepository repository.ProductLoader) *ProductServices {
	return &ProductServices{productRepository: productRepository}
}
