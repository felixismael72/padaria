package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"padaria/src/app/api/endpoints/dto/request"
	"padaria/src/app/api/endpoints/dto/response"
	"padaria/src/core/interfaces/primary"
)

type ProductHandlers struct {
	productService primary.ProductManager
}

func (handler ProductHandlers) PostProduct(c echo.Context) error {
	var product request.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Algo está incompatível na sua requisição.",
			http.StatusBadRequest,
		),
		)
	}

	productID, registerErr := handler.productService.RegisterProduct(*product.ToDomain())
	if registerErr != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusInternalServerError,
		),
		)
	}

	return c.JSON(http.StatusCreated, &response.Created{ID: productID})
}

func NewProductHandlers(productService primary.ProductManager) *ProductHandlers {
	return &ProductHandlers{productService: productService}
}
