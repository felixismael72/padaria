package router

import (
	"padaria/src/app/api/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadProductRoutes(api *echo.Group) {
	productGroup := api.Group("/product")

	productHandlers := dicontainer.GetProductHandlers()

	productGroup.POST("/new", productHandlers.PostProduct)
	productGroup.GET("", productHandlers.GetProducts)
	productGroup.PUT("/:productID/edit", productHandlers.PutProduct)
}
