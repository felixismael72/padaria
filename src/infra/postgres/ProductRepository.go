package postgres

import (
	"log"
	"padaria/src/core/domain"
	"padaria/src/core/interfaces/repository"
)

var _ repository.ProductLoader = (*ProductRepository)(nil)

type ProductRepository struct {
	connectorManager
}

func (repo ProductRepository) InsertProduct(product domain.Product) (int, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Print(err)
		return -1, err
	}
	defer repo.closeConnection(conn)

	query := `insert into product(name, code, price, expiration_date) 
			  values($1, $2, $3, $4) returning id;`

	var productID int
	err = conn.Get(
		&productID,
		query,
		product.Name(),
		product.Code(),
		product.Price(),
		product.ExpirationDate(),
	)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return productID, nil
}

func NewProductRepository(manager connectorManager) *ProductRepository {
	return &ProductRepository{manager}
}
