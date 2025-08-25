package service

import (
	"dacode/OneDrive/Desktop/2706NewProj/internal/pkg/cart/model"
	"errors"
	"fmt"
)

type CartRepository interface {
	AddProduct(userID int64, sku int64, count uint16) error
	DeleteProduct(userID int64, sku int64) error
	ClearCart(userID int64) error
	GetCart(userID int64) (*model.Cart, error)
}

type ProductClient interface {
	GetProductInfo(sku int64) (name string, price uint32, err error)
}

type CartService struct {
	repository    CartRepository
	productClient ProductClient
}

func NewService(repository CartRepository, productClient ProductClient) *CartService {
	return &CartService{
		repository:    repository,
		productClient: productClient,
	}
}

func (s *CartService) AddProduct(userID int64, sku int64, count uint16) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	return s.repository.AddProduct(userID, sku, count)
}

func (s *CartService) DeleteProduct(userID int64, sku int64, count uint16) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	return s.repository.DeleteProduct(userID, sku, count)
}
func (s *CartService) ClearCart(userID int64) error {
	if userID < 1 {
		return errors.New("userID and sku must be defined")
	}
	return s.repository.ClearCart(userID)
}

func (s *CartService) GetCart(userID int64) (*model.Cart, error) {
	if userID < 1 {
		return nil, errors.New("userID must be defined")
	}

	cart, err := s.repository.GetCart(userID)

	if err != nil {
		return nil, err
	}
	for i := range cart.Item {
		name, price, err := s.productClient.GetProductInfo(cart.Item[i].SkuID)

		if err != nil {
			return nil, fmt.Errorf("failed to get product info: %w", err)
		}

		cart.Item[i].Name = name
		cart.Item[i].Price = price

	}

	var totalPrice uint32

	for _, item := range cart.Item {
		totalPrice += item.Price * uint32(item.Count)
	}

	cart.TotalPrice = totalPrice

	return cart, nil

}
