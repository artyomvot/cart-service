package server

import (
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(cart)
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
