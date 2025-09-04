package server

import (
	"2706NewProj/internal/pkg/cart/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type PProduct struct {
	SkuID int64  `json:"sku_id"`
	Count uint16 `json:"count"`
}

type DeleteProductRequest struct {
	UserID int64              `json:"user_id"`
	Item   map[int64]PProduct `json:"item"`
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	rawUID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID format")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
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
			log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	if sku < 1 || userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "User ID and SKU must be positive integers")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	var deleteRequest DeleteProductRequest
	err = json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid JSON format in request body")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	if deleteRequest.UserID < 1 || len(deleteRequest.Item) == 0 || deleteRequest.UserID != userID {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID or empty items in request")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}
		return
	}

	delItem := make(map[int64]*model.Product, len(deleteRequest.Item))

	for skuID, product := range deleteRequest.Item {
		delItem[skuID] = &model.Product{
			Count: product.Count,
			SkuID: product.SkuID,
		}
	}

	for _, v := range delItem {
		err := s.cartService.DeleteProduct(userID, v.SkuID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to delete product: %s\"}", err.Error())
			if errOut != nil {
				log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Products deleted from cart successfully")
	if errOut != nil {
		log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
		return
	}
}
