package handlers

import (
	"net/http"
	"strconv"

	"padaria/src/app/api/endpoints/dto/request"
	"padaria/src/app/api/endpoints/dto/response"
	"padaria/src/core/interfaces/primary"

	"github.com/labstack/echo/v4"
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

func (handler ProductHandlers) GetProducts(c echo.Context) error {
	products, err := handler.productService.ListProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusInternalServerError,
		),
		)
	}

	var productDTOList []response.Product
	for _, product := range products {
		productDTOList = append(productDTOList, *response.NewProduct(product))
	}

	return c.JSON(http.StatusOK, productDTOList)
}

func (handler ProductHandlers) PutProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"O id desse produto não pôde ser processado.",
			http.StatusBadRequest,
		),
		)
	}

	var productDTO request.Product
	if err := c.Bind(&productDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Algo está incompatível na sua requisição.",
			http.StatusBadRequest,
		),
		)
	}

	product := productDTO.ToDomainWithID(productID)

	err = handler.productService.EditProduct(*product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusInternalServerError,
		),
		)
	}

	return c.NoContent(http.StatusNoContent)
}

func NewProductHandlers(productService primary.ProductManager) *ProductHandlers {
	return &ProductHandlers{productService: productService}
}
