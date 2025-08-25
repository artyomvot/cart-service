package repository

import (
	"dacode/OneDrive/Desktop/2706NewProj/internal/pkg/cart/model"
	"errors"
)

type Storage = map[int64]model.Cart

type Repository struct {
	storage Storage
}

func NewCartRepository(capacity int) *Repository {
	return &Repository{storage: make(Storage, capacity)}
}

func (r *Repository) AddProduct(userID int64, sku int64, count uint16) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	cart, ok := r.storage[userID]
	if !ok {
		cart = model.Cart{
			UserID: userID,
			Item:   []model.Product{},
		}
	}
	newProduct := model.Product{
		SkuID: sku,
		Count: count,
	}
	cart.Item = append(cart.Item, newProduct)
	r.storage[userID] = cart
	return nil
}

func (r *Repository) DeleteProduct(userID int64, sku int64) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	cart, ok := r.storage[userID]
	if !ok {
		return errors.New("Unable to remove product from non-existent cart")
	}
	for i, _ := range cart.Item {
		if sku == cart.Item[i].SkuID {
			cart.Item = append(cart.Item[:i], cart.Item[i+1:]...)
			return nil
		}
	}
	return errors.New("Product not found")
}

func (r *Repository) ClearCart(userID int64) error {
	if userID < 1 {
		return errors.New("userID must be defined")
	}

	cart, ok := r.storage[userID]
	if !ok {
		return errors.New("the cart must exist")
	}

	cart.Item = []model.Product{}
	return nil

}
