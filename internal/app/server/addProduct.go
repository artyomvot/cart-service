package server

import (
	"ProductCartService/internal/pkg/cart/productClient"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AddProductRequest struct {
	Count uint16 `json:"count"`
}

func (s *Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawID, 10, 64)
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

	rawID = r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawID, 10, 64)
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

	if addRequest.Count < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid Count")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	err = s.cartService.AddProduct(userID, sku, addRequest.Count)
	if err != nil {
		if errors.Is(err, productClient.ErrNotFound) {
			w.WriteHeader(http.StatusPreconditionFailed)
			w.Header().Set("Content-Type", "application/json")
			_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to add product to cart: %s\"}", err.Error())
			if errOut != nil {
				log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to add product to cart: %s\"}", err.Error())
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Products added to cart successfully")
	if errOut != nil {
		log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
		return
	}
}
