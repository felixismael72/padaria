package services

import (
	"errors"
	"testing"
	"time"

	"padaria/src/core/domain"
	repoMock "padaria/src/core/interfaces/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterProduct(t *testing.T) {
	caseManager := registerProductCases{}

	t.Run("Test RegisterProduct when it has a happy path", caseManager.testWhenItHasAHappyPath)
	t.Run("Test RegisterProduct when it receives an error", caseManager.testWhenItReceivesAnError)
}

type registerProductCases struct{}

func (registerProductCases) testWhenItHasAHappyPath(t *testing.T) {
	controller := gomock.NewController(t)
	productRepository := repoMock.NewMockProductLoader(controller)
	productService := NewProductServices(productRepository)
	product := domain.NewProduct(0, "ProdutoTest", "0000", 0.1, time.Now().Add(30*time.Second))
	expectedID := 1
	productRepository.EXPECT().InsertProduct(*product).Return(expectedID, nil).MaxTimes(1)

	actualID, actualErr := productService.RegisterProduct(*product)

	assert.Equal(t, expectedID, actualID)
	assert.Nil(t, actualErr)
}

func (registerProductCases) testWhenItReceivesAnError(t *testing.T) {
	controller := gomock.NewController(t)
	productRepository := repoMock.NewMockProductLoader(controller)
	productService := NewProductServices(productRepository)
	product := domain.NewProduct(0, "ProdutoTest", "0000", 0.1, time.Now().Add(30*time.Second))
	expectedID := -1
	expectedErr := errors.New("the database connection could not be established")
	productRepository.EXPECT().InsertProduct(*product).Return(expectedID, expectedErr).MaxTimes(1)

	actualID, actualErr := productService.RegisterProduct(*product)

	assert.Equal(t, expectedID, actualID)
	assert.Equal(t, expectedErr, actualErr)
}
