package router

import (
	"github.com/labstack/echo/v4"
	"padaria/src/app/api/dicontainer"
)

func loadProductRoutes(api *echo.Group) {
	productGroup := api.Group("/product")

	productHandlers := dicontainer.GetProductHandlers()

	productGroup.POST("/new", productHandlers.PostProduct)
}
