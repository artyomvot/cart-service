package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawID, 10, 64)
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

	rawID = r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawID, 10, 64)
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
	err = s.cartService.DeleteProduct(userID, sku)
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Products deleted from cart successfully")
	if errOut != nil {
		log.Printf("DELETE /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
		return
	}
}
