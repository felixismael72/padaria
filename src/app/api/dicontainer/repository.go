package dicontainer

import (
	"padaria/src/core/interfaces/repository"
	"padaria/src/infra/postgres"
)

func GetProductRepository() repository.ProductLoader {
	return postgres.NewProductRepository(GetPSQLConnector())
}

func GetPSQLConnector() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
