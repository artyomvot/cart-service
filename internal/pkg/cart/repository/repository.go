package repository

import (
	"ProductCartService/internal/pkg/cart/model"
	"errors"
)

type Storage = map[int64]*model.Cart // ключ - юзерАйДи

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
	if count < 1 {
		return errors.New("count must be more than 0")
	}
	cart, ok := r.storage[userID]
	if !ok {
		newcart := model.Cart{
			UserID: userID,
			Item:   make(map[int64]*model.Product),
		}
		r.storage[userID] = &newcart
		cart = r.storage[userID]
	}
	_, ok = cart.Item[sku]
	if ok {
		cart.Item[sku].Count += count
		return nil
	}

	newProduct := model.Product{
		SkuID: sku,
		Count: count,
	}
	cart.Item[sku] = &newProduct
	return nil
}

func (r *Repository) DeleteProduct(userID int64, sku int64) error {
	if userID < 1 || sku < 1 {
		return errors.New("userID and sku must be defined")
	}
	cart, ok := r.storage[userID]
	if !ok {
		return nil
	}

	_, ok = cart.Item[sku]
	if ok {
		delete(cart.Item, sku)
		return nil
	}
	return nil
}

func (r *Repository) ClearCart(userID int64) error {
	if userID < 1 {
		return errors.New("userID must be defined")
	}

	cart, ok := r.storage[userID]
	if !ok {
		return nil
	}

	cart.Item = make(map[int64]*model.Product)
	return nil

}

func (r *Repository) GetCart(userID int64) (*model.Cart, error) {
	if userID < 1 {
		return nil, errors.New("userID must be defined")
	}

	cart, ok := r.storage[userID]

	if !ok {
		return nil, nil
	}

	return cart, nil

}
