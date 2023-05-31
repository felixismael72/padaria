package dicontainer

import (
	"padaria/src/core/interfaces/primary"
	"padaria/src/core/services"
)

func GetProductServices() primary.ProductManager {
	return services.NewProductServices(GetProductRepository())
}
