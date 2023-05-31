package dicontainer

import "padaria/src/app/api/endpoints/handlers"

func GetProductHandlers() *handlers.ProductHandlers {
	return handlers.NewProductHandlers(GetProductServices())
}
