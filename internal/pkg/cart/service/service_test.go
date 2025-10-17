package service

import (
	"ProductCartService/internal/pkg/cart/model"
	"ProductCartService/internal/pkg/cart/productClient"

	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestService_AddProduct_Success(t *testing.T) {
	type data struct {
		name       string
		userID     int64
		SkuID      int64
		Count      uint16
		clientItem *productClient.Item
		wantErr    error
	}
	testData := []data{
		{
			name:   "valid add item",
			userID: 123,
			SkuID:  123456,
			Count:  2,
			clientItem: &productClient.Item{
				Name:  "Книга",
				Price: 300,
			},
			wantErr: nil,
		},
	}

	ctrl := minimock.NewController(t)
	productMock := NewProductClientMock(ctrl)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, productMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			productMock.GetProductInfoMock.Expect(tt.SkuID).Return(tt.clientItem, nil)
			repoMock.AddProductMock.Expect(tt.userID, tt.SkuID, tt.Count).Return(nil)
			err := serviceMock.AddProduct(tt.userID, tt.SkuID, tt.Count)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}

}

func TestService_AddProduct_Fail(t *testing.T) {
	type data struct {
		name       string
		userID     int64
		SkuID      int64
		Count      uint16
		clientItem *productClient.Item
		wantErr    error
	}
	testData := []data{
		{
			name:       "Product not found",
			userID:     123,
			SkuID:      234567,
			Count:      2,
			clientItem: nil,
			wantErr:    errors.New("product not found in external service"),
		},
		{
			name:   "invalid userID",
			userID: -123,
			SkuID:  123456,
			Count:  2,
			clientItem: &productClient.Item{
				Name:  "Книга",
				Price: 300,
			},
			wantErr: errors.New("userID and sku must be defined"),
		},
		{
			name:   "invalid sku",
			userID: 123,
			SkuID:  -100,
			Count:  2,
			clientItem: &productClient.Item{
				Name:  "Книга",
				Price: 300,
			},
			wantErr: errors.New("userID and sku must be defined"),
		},
	}

	ctrl := minimock.NewController(t)
	productMock := NewProductClientMock(ctrl)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, productMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Product not found" {
				productMock.GetProductInfoMock.Expect(tt.SkuID).Return(nil, errors.New("product not found in external service"))
			}
			err := serviceMock.AddProduct(tt.userID, tt.SkuID, tt.Count)
			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestService_DeleteProduct_Success(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		sku     int64
		wantErr error
	}
	testData := []data{
		{
			name:    "valid add item",
			userID:  123,
			sku:     123456,
			wantErr: nil,
		},
		{
			name:    "delete non-existing product",
			userID:  123,
			sku:     999999,
			wantErr: nil,
		},
	}
	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, nil)
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repoMock.DeleteProductMock.Expect(tt.userID, tt.sku).Return(nil)
			err := serviceMock.DeleteProduct(tt.userID, tt.sku)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}

}

func TestService_DeleteProduct_Fail(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		sku     int64
		wantErr error
	}
	testData := []data{
		{
			name:    "invalid userID",
			userID:  -123,
			sku:     123456,
			wantErr: errors.New("userID and sku must be defined"),
		},
		{
			name:    "invalid sku",
			userID:  123,
			sku:     -123456,
			wantErr: errors.New("userID and sku must be defined"),
		},
	}

	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, nil)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceMock.DeleteProduct(tt.userID, tt.sku)
			require.EqualError(t, err, tt.wantErr.Error())
		})
	}

}

func TestService_ClearCart_Success(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		wantErr error
	}

	testData := []data{
		{
			name:    "valid clear cart",
			userID:  123,
			wantErr: nil,
		},
	}

	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, nil)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repoMock.ClearCartMock.Expect(tt.userID).Return(nil)
			err := serviceMock.ClearCart(tt.userID)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestService_ClearCart_Fail(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		wantErr error
	}

	testData := []data{
		{
			name:    "invalid userID",
			userID:  -123,
			wantErr: errors.New("userID must be defined"),
		},
	}
	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	serviceMock := NewService(repoMock, nil)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceMock.ClearCart(tt.userID)
			require.EqualError(t, err, tt.wantErr.Error())
		})
	}
}

func TestService_GetCart_Success(t *testing.T) {
	type data struct {
		name        string
		userID      int64
		repoCart    *model.Cart
		clientItems map[int64]*productClient.Item
		wantCart    *model.Cart
		wantErr     error
	}
	testData := []data{
		{
			name:   "valid get cart with 1 item",
			userID: 123,
			repoCart: &model.Cart{
				UserID: 123,
				Item: map[int64]*model.Product{
					123456: {SkuID: 123456, Count: 2},
				},
			},
			clientItems: map[int64]*productClient.Item{
				123456: {Name: "Книга", Price: 500},
			},
			wantCart: &model.Cart{
				UserID: 123,
				Item: map[int64]*model.Product{
					123456: {SkuID: 123456, Count: 2, Name: "Книга", Price: 500},
				},
				TotalPrice: 1000,
			},
			wantErr: nil,
		},
		{
			name:   "empty cart",
			userID: 123,
			repoCart: &model.Cart{
				UserID: 123,
				Item:   map[int64]*model.Product{},
			},
			clientItems: nil,
			wantCart:    nil,
			wantErr:     nil,
		},
	}

	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	productMock := NewProductClientMock(ctrl)
	serviceMock := NewService(repoMock, productMock)
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repoMock.GetCartMock.Expect(tt.userID).Return(tt.repoCart, nil)
			if tt.clientItems != nil {
				for sku, item := range tt.clientItems {
					productMock.GetProductInfoMock.Expect(sku).Return(item, nil)
				}
			}
			wantCart, err := serviceMock.GetCart(tt.userID)

			require.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				require.Equal(t, wantCart, tt.wantCart)
			}
		})
	}

}

func TestService_GetCart_Fail(t *testing.T) {
	type data struct {
		name        string
		userID      int64
		repoCart    *model.Cart
		clientItems map[int64]*productClient.Item
		wantCart    *model.Cart
		wantErr     error
	}
	testData := []data{
		{
			name:        "invalid userID",
			userID:      -123,
			repoCart:    nil,
			clientItems: nil,
			wantCart:    nil,
			wantErr:     errors.New("userID must be defined"),
		},
		{
			name:        "failed to get product from repository",
			userID:      123,
			repoCart:    nil,
			clientItems: nil,
			wantCart:    nil,
			wantErr:     errors.New("failed to get product from repository"),
		},
		{
			name:   "failed to get product info",
			userID: 123,
			repoCart: &model.Cart{
				UserID: 123,
				Item: map[int64]*model.Product{
					123456: {SkuID: 123456, Count: 2},
				},
			},
			clientItems: nil,
			wantCart:    nil,
			wantErr:     errors.New("failed to get product info"),
		},
	}
	ctrl := minimock.NewController(t)
	repoMock := NewCartRepositoryMock(ctrl)
	productMock := NewProductClientMock(ctrl)
	serviceMock := NewService(repoMock, productMock)
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "invalid userID" {
				_, err := serviceMock.GetCart(tt.userID)
				require.EqualError(t, err, tt.wantErr.Error())
			} else if tt.name == "failed to get product from repository" {
				repoMock.GetCartMock.Expect(tt.userID).Return(nil, errors.New("failed to get product from repository"))
				_, err := serviceMock.GetCart(tt.userID)
				require.EqualError(t, err, tt.wantErr.Error())
			} else if tt.name == "failed to get product info" {
				repoMock.GetCartMock.Expect(tt.userID).Return(tt.repoCart, nil)
				for sku := range tt.repoCart.Item {
					productMock.GetProductInfoMock.Expect(sku).Return(nil, errors.New("failed to get product info"))
				}
				_, err := serviceMock.GetCart(tt.userID)
				require.EqualError(t, err, tt.wantErr.Error())
			}

		})
	}
}
