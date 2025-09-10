package server

import (
	"ProductCartService/internal/pkg/cart/model"
)

type CartService interface {
	AddProduct(userID int64, sku int64, count uint16) error
	DeleteProduct(userID int64, sku int64) error
	ClearCart(userID int64) error
	GetCart(userID int64) (*model.Cart, error)
}

type Server struct {
	cartService CartService
}

func New(cartService CartService) *Server {
	return &Server{cartService: cartService}
}
