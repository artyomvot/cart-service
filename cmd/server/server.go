package server

import (
	"dacode/OneDrive/Desktop/2706NewProj/internal/app/server"
	"dacode/OneDrive/Desktop/2706NewProj/internal/pkg/cart/repository"
	"dacode/OneDrive/Desktop/2706NewProj/internal/pkg/cart/service"
	"net/http"
)

func main() {

	cartRepository := repository.NewCartRepository(100)
	cartService := service.NewService(cartRepository)

	server := server.New(cartService)

	http.ListenAndServe(":0808", nil)

}
