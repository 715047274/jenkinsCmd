package App

import (
	"github.com/gin/demo/internal/registry"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock implementation of the CartService
type MockCartService struct {
	mock.Mock
}

//func (m *MockCartService) AddToCart(userID, productID int64, quantity int) error {
//	args := m.Called(userID, productID, quantity)
//	return args.Error(0)
//}
//
//func (m *MockCartService) GetCart(userID int64) ([]domain.CartItem, error) {
//	args := m.Called(userID)
//	return args.Get(0).([]domain.CartItem), args.Error(1)
//}

func TestDynamicServiceRegistry(t *testing.T) {
	// Setup mock service
	mockCartService := new(MockCartService)
	mockCartService.On("AddToCart", int64(1), int64(100), 2).Return(nil)

	// Register mock service in a new service registry
	serviceRegistry := registry.NewServiceRegistry()
	serviceRegistry.RegisterService("CartService", mockCartService)

	// Retrieve CartService dynamically
	//serviceInterface, err := serviceRegistry.GetService("CartService")
	//assert.NoError(t, err)

	// Type assert to CartService
	//cartService, ok := serviceInterface.(application.CartService)
	//assert.True(t, ok)
	//
	//// Call AddToCart and verify the mock interaction
	//err = cartService.AddToCart(1, 100, 2)
	//assert.NoError(t, err)
	//mockCartService.AssertCalled(t, "AddToCart", int64(1), int64(100), 2)
}
