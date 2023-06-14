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

func (service ProductServices) ListProducts() ([]domain.Product, error) {
	products, err := service.productRepository.SelectProducts()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return products, nil
}

func (service ProductServices) EditProduct(product domain.Product) error {
	err := service.productRepository.UpdateProduct(product)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (service ProductServices) RemoveProduct(productID int) error {
	err := service.productRepository.DeleteProduct(productID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func NewProductServices(productRepository repository.ProductLoader) *ProductServices {
	return &ProductServices{productRepository: productRepository}
}
