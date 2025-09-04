package main

import (
	"net/http"

	"2706NewProj/internal/app/server"
	"2706NewProj/internal/pkg/cart/productClient"
	"2706NewProj/internal/pkg/cart/repository"
	"2706NewProj/internal/pkg/cart/service"
)

func main() {
	cartRepository := repository.NewCartRepository(100)

	productClient := productClient.New(
		"http://route256.pavl.uk:8080",
		"testtoken",
	)

	cartService := service.NewService(cartRepository, productClient)
	srv := server.New(cartService)

	http.HandleFunc("POST /user/{user_id}/cart/{sku_id}", srv.AddProduct)
	http.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", srv.DeleteProduct)
	http.HandleFunc("DELETE /user/{user_id}/cart", srv.ClearCart)
	http.HandleFunc("GET /user/{user_id}/cart", srv.GetCart)

	http.ListenAndServe(":8082", nil)
}
