package server

import "fmt"
// пока хз
type CartService interface {
	AddProduct()
	DeleteProduct()
	ClearCart()
	GetCart()
}

type Server struct{
	cartService CartService
}

func New(cartService CartService) *Server{
	return &Server{cartService: cartService}
}