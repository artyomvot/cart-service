package server

import (
	"2706NewProj/internal/pkg/cart/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	SkuID int64  `json:"sku_id"`
	Count uint16 `json:"count"`
}

type AddProductRequest struct {
	UserID int64             `json:"user_id"`
	Item   map[int64]Product `json:"item"`
}

func (s *Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	rawUID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID format")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	rawSID := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawSID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid SKU ID format")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	if sku < 1 || userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "User ID and SKU must be positive integers")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	var addRequest AddProductRequest
	err = json.NewDecoder(r.Body).Decode(&addRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid JSON format in request body")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	if addRequest.UserID < 1 || len(addRequest.Item) == 0 || addRequest.UserID != userID {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID or empty items in request")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	inputItem := make(map[int64]*model.Product, len(addRequest.Item))

	for skuID, item := range addRequest.Item {
		inputItem[skuID] = &model.Product{
			Count: item.Count,
			SkuID: item.SkuID,
		}
	}

	for _, v := range inputItem {
		err := s.cartService.AddProduct(userID, v.SkuID, v.Count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to add product to cart: %s\"}", err.Error())
			if errOut != nil {
				log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Products added to cart successfully")
	if errOut != nil {
		log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
		return
	}
}
