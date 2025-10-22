package server

import (
	"ProductCartService/internal/pkg/cart/productClient"
	"ProductCartService/internal/pkg/cart/repository"
	"ProductCartService/internal/pkg/cart/service"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetCartSuite struct {
	suite.Suite
	cartService *service.CartService
}

func (g *GetCartSuite) SetupSuite() {
	testStorage := repository.NewCartRepository(10)
	testProductService := productClient.New("http://route256.pavl.uk:8080", "testtoken")
	g.cartService = service.NewService(testStorage, testProductService)
}

func (g *GetCartSuite) TestGetCart() {
	userID := int64(123)
	sku := int64(773297411)
	count := uint16(2)
	err := g.cartService.AddProduct(userID, sku, count)
	require.NoError(g.T(), err)
	cart, err := g.cartService.GetCart(userID)
	require.NoError(g.T(), err)
	require.NotNil(g.T(), cart)
	require.Contains(g.T(), cart.Item, sku)
	require.Equal(g.T(), count, cart.Item[sku].Count)
}

func TestGetCartSuite(t *testing.T) {
	suite.Run(t, new(GetCartSuite))
}
