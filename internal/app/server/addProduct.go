package server

import (
	"dacode/OneDrive/Desktop/2706NewProj/internal/pkg/cart/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	SkuID int64  `json:"sku_id"`
	Count uint16 `json:"count"`
}

type AddProductRequest struct {
	UserID int64     `json:"user_id"`
	Item   []Product `json:"item"`
}

func (s *Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	rawUID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUID, 10, 64)
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
	rawSID := r.PathValue("sku_id")
	sku, erro := strconv.ParseInt(rawSID, 10, 64)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}

	if sku < 1 || userID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "sku must be more than 0")
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}

	body, err := io.ReadAll(r.Body)

	var addRequest AddProductRequest

	err = json.Unmarshal(body, &addRequest)

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

	if addRequest.UserID < 1 || len(addRequest.Item) == 0 || addRequest.UserID != userID {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/<user_id>/cart/<sku_id> out failed: %s", errOut.Error())
			return
		}

		return
	}

	inputProduct := model.Cart{
		UserID: addRequest.UserID,
		Item:   addRequest.Item,
	}

}
