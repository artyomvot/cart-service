package repository

import (
	"2706NewProj/internal/pkg/cart/model"
	"errors"
)

type Storage = map[int64]model.Cart // ключ - юзерАйДи

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
		cart = model.Cart{
			UserID: userID,
			Item:   make(map[int64]*model.Product),
		}
	}

	for i := range cart.Item {
		if cart.Item[i].SkuID == sku {
			cart.Item[i].Count += count
			r.storage[userID] = cart
		}
		return nil
	}

	newProduct := model.Product{
		SkuID: sku,
		Count: count,
	}
	cart.Item[sku] = &newProduct
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
			delete(cart.Item, i)
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

	cart.Item = make(map[int64]*model.Product)
	return nil

}

func (r *Repository) GetCart(userID int64) (*model.Cart, error) {
	if userID < 1 {
		return nil, errors.New("userID must be defined")
	}

	cart, ok := r.storage[userID]

	if !ok {
		return nil, errors.New("cart not found")
	}

	return &cart, nil

}
