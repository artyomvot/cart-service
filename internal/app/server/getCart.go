package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type GetCartProduct struct {
	SkuID int64  `json:"sku_id"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type GetCartResponse struct {
	TotalPrice uint32 `json:"total_price"`
	Product    []GetCartProduct
}

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")

	userID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID format")
		if errOut != nil {
			log.Printf("GET /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	if userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "User ID must be positive integer")
		if errOut != nil {
			log.Printf("GET /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	cart, err := s.cartService.GetCart(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to get cart: %s\"}", err.Error())
		if errOut != nil {
			log.Printf("GET /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	res := GetCartResponse{
		TotalPrice: cart.TotalPrice,
		Product:    make([]GetCartProduct, 0, len(cart.Item)),
	}

	for _, item := range cart.Item {
		res.Product = append(res.Product, GetCartProduct{
			SkuID: item.SkuID,
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to encode cart to JSON\"}")
		if errOut != nil {
			log.Printf("GET /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("GET /user/<user_id>/cart out failed: %s", err.Error())
		return
	}
}
