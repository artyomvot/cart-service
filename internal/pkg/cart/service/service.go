package service

import (
	"ProductCartService/internal/pkg/cart/model"
	"ProductCartService/internal/pkg/cart/productClient"
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
	GetProductInfo(sku int64) (*productClient.Item, error)
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
	_, err := s.productClient.GetProductInfo(sku)
	if err != nil {
		return fmt.Errorf("product %d not found in external service: %w", sku, err)
	}
	return s.repository.AddProduct(userID, sku, count)
}

func (s *CartService) DeleteProduct(userID int64, sku int64) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	return s.repository.DeleteProduct(userID, sku)
}
func (s *CartService) ClearCart(userID int64) error {
	if userID < 1 {
		return errors.New("userID must be defined")
	}
	return s.repository.ClearCart(userID)
}

func (s *CartService) GetCart(userID int64) (*model.Cart, error) {
	if userID < 1 {
		return nil, errors.New("userID must be defined")
	}

	cart, err := s.repository.GetCart(userID)
	if err != nil {
		return nil, errors.New("failed to get product from repository")
	}
	if cart == nil || len(cart.Item) == 0 {
		return nil, nil
	}

	for skuID := range cart.Item {
		var item *productClient.Item
		item, err := s.productClient.GetProductInfo(cart.Item[skuID].SkuID)

		if err != nil {
			return nil, errors.New("failed to get product info")
		}

		cart.Item[skuID].Name = item.Name
		cart.Item[skuID].Price = item.Price

	}

	var totalPrice uint32

	for _, item := range cart.Item {
		totalPrice += item.Price * uint32(item.Count)
	}

	cart.TotalPrice = totalPrice

	return cart, nil

}
