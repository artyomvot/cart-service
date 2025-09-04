package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) ClearCart(w http.ResponseWriter, r *http.Request) {
	rawUID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "Invalid user ID format")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	if userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "User ID must be positive integer")
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	err = s.cartService.ClearCart(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"Failed to clear cart: %s\"}", err.Error())
		if errOut != nil {
			log.Printf("DELETE /user/<user_id>/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"message": "Cart cleared successfully"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal success response: %s", err)
		return
	}
	_, errOut := w.Write(jsonResp)
	if errOut != nil {
		log.Printf("Failed to write success response: %s", errOut.Error())
	}
}
