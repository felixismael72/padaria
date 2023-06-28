package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"padaria/src/core/domain"
	serviceMock "padaria/src/core/interfaces/primary/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	postProductURL = "/api/product/new"
)

func TestPostProduct(t *testing.T) {
	caseManager := postProductCases{}

	t.Run("Test PostProduct when it has a happy path", caseManager.testWhenItHasAHappyPath)
	t.Run("Test PostProduct when it cannot process the product", caseManager.testWhenItCannotProcessTheProduct)
	t.Run("Test PostProduct when it receives a connection error", caseManager.testWhenItReceivesAConnectionError)
}

type postProductCases struct{}

func (postProductCases) testWhenItHasAHappyPath(t *testing.T) {
	controller := gomock.NewController(t)
	productService := serviceMock.NewMockProductManager(controller)
	product := domain.NewProduct(0, "ProdutoTest", "0000", 0.1, time.Time{})
	expectedID := 1
	productService.EXPECT().RegisterProduct(*product).Return(expectedID, nil).MaxTimes(1)
	productHandlers := NewProductHandlers(productService)
	requestBody := fmt.Sprintf(`
			{
				"name": "%s",
				"code": "%s",
				"price": %f
			}
		`,
		product.Name(),
		product.Code(),
		product.Price(),
	)
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	request.Header.Add(echo.HeaderContentType, "application/json")
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedJSON := fmt.Sprintf(`{"id":%d}`, expectedID)
	expectedStatusCode := http.StatusCreated

	_ = productHandlers.PostProduct(context)

	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}

func (postProductCases) testWhenItCannotProcessTheProduct(t *testing.T) {
	controller := gomock.NewController(t)
	productService := serviceMock.NewMockProductManager(controller)
	productHandlers := NewProductHandlers(productService)
	requestBody := `{}`
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedJSON := fmt.Sprintf(`
			{
				"msg": "%s",
				"statusCode": %d
			}
		`,
		"Algo está incompatível na sua requisição.",
		http.StatusBadRequest,
	)
	expectedStatusCode := http.StatusBadRequest

	_ = productHandlers.PostProduct(context)

	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}

func (postProductCases) testWhenItReceivesAConnectionError(t *testing.T) {
	controller := gomock.NewController(t)
	productService := serviceMock.NewMockProductManager(controller)
	err := errors.New("connection error")
	productService.EXPECT().RegisterProduct(gomock.Any()).Return(-1, err).MaxTimes(1)
	productHandlers := NewProductHandlers(productService)
	requestBody := `{}`
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	request.Header.Add(echo.HeaderContentType, "application/json")
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedJSON := fmt.Sprintf(`
			{
				"msg": "%s",
				"statusCode": %d
			}
		`,
		"Oops! Parece que o serviço de dados está indisponível.",
		http.StatusInternalServerError,
	)
	expectedStatusCode := http.StatusInternalServerError

	_ = productHandlers.PostProduct(context)

	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}
