package server

import (
	"ProductCartService/internal/pkg/cart/productClient"
	"ProductCartService/internal/pkg/cart/repository"
	"ProductCartService/internal/pkg/cart/service"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeleteItemSuite struct {
	suite.Suite
	cartService *service.CartService
}

func (d *DeleteItemSuite) SetupSuite() {
	storage := repository.NewCartRepository(10)
	productClient := productClient.New("http://route256.pavl.uk:8080", "testtoken")
	d.cartService = service.NewService(storage, productClient)

}

func (d *DeleteItemSuite) TestDeleteItem() {
	userID := int64(123)
	sku := int64(773297411)
	count := uint16(2)
	err := d.cartService.AddProduct(userID, sku, count)
	require.NoError(d.T(), err)
	err = d.cartService.DeleteProduct(userID, sku)
	require.NoError(d.T(), err)
	cart, err := d.cartService.GetCart(userID)
	require.NoError(d.T(), err)

	require.Nil(d.T(), cart)

}

func TestDeleteItemSuite(t *testing.T) {
	suite.Run(t, new(DeleteItemSuite))
}
