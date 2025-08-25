package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")

	userID, err := strconv.ParseInt(rawID, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}

	if userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "sku must be more than 0")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}

	cart, err := s.cartService.GetCart(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "sku must be more than 0")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
