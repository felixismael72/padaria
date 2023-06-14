package postgres

import (
	"log"

	"padaria/src/core/domain"
	"padaria/src/core/interfaces/repository"
	"padaria/src/infra/postgres/dto"
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

func (repo ProductRepository) SelectProducts() ([]domain.Product, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer repo.closeConnection(conn)

	query := `select 
							id 							as product_id,
							name 						as product_name,
							code 						as product_code,
							price 					as product_price,
							expiration_date as product_expiration_date
						from product;`

	var productDTOList []dto.Product
	err = conn.Select(&productDTOList, query)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var products []domain.Product
	for _, productDTO := range productDTOList {
		products = append(products, *productDTO.ToDomain())
	}

	return products, nil
}

func NewProductRepository(manager connectorManager) *ProductRepository {
	return &ProductRepository{manager}
}
