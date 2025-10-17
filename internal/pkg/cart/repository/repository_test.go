package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddItem(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		sku     int64
		count   uint16
		wantErr error
	}
	testData := []data{
		{
			name:    "Success test",
			userID:  123,
			sku:     123456,
			count:   2,
			wantErr: nil,
		},
		{
			name:    "invalid UserID",
			userID:  -3,
			sku:     123456,
			count:   2,
			wantErr: errors.New("userID and sku must be defined"),
		},
		{
			name:    "invalid SKU",
			userID:  123,
			sku:     -256987,
			count:   2,
			wantErr: errors.New("userID and sku must be defined"),
		},
		{
			name:    "invalid Count",
			userID:  123,
			sku:     123456,
			count:   0,
			wantErr: errors.New("count must be more than 0"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCartRepository(10)
			err := repo.AddProduct(tt.userID, tt.sku, tt.count)
			if tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestDeleteItem(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		sku     int64
		wantErr error
	}
	testData := []data{
		{
			name:    "success delete",
			userID:  123,
			sku:     123456,
			wantErr: nil,
		},
		{
			name:    "invalid userID",
			userID:  -123,
			sku:     123456,
			wantErr: errors.New("userID and sku must be defined"),
		},
		{
			name:    "invalid SKU",
			userID:  123,
			sku:     -123456,
			wantErr: errors.New("userID and sku must be defined"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCartRepository(10)
			err := repo.DeleteProduct(tt.userID, tt.sku)

			if err != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClearCart(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		wantErr error
	}
	testData := []data{
		{
			name:    "Success clear",
			userID:  123,
			wantErr: nil,
		},
		{
			name:    "invalid userID",
			userID:  -123,
			wantErr: errors.New("userID must be defined"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCartRepository(10)
			err := repo.ClearCart(tt.userID)
			if err != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestGetCart(t *testing.T) {
	type data struct {
		name    string
		userID  int64
		wantErr error
	}
	testData := []data{
		{
			name:    "Success GetCart",
			userID:  123,
			wantErr: nil,
		},
		{
			name:    "invalid userID",
			userID:  -123,
			wantErr: errors.New("userID must be defined"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCartRepository(10)
			_, err := repo.GetCart(tt.userID)
			if err != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

		})
	}

}
