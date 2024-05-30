package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/helpers"
	"github.com/matsuboshi/ramengo/internal/models"
)

type OrderInput struct {
	ProteinId string `json:"proteinId"`
	BrothId   string `json:"brothId"`
}

func PostOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.CustomError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var orderInput OrderInput

	if err := json.NewDecoder(r.Body).Decode(&orderInput); err != nil {
		errorMessage := fmt.Sprintf("error decoding order input: %v", err)
		helpers.CustomError(w, errorMessage, http.StatusBadRequest)
		return
	}

	secretKey := r.Header.Get("x-api-key")
	if secretKey == "" {
		helpers.CustomError(w, "x-api-key header missing", http.StatusForbidden)
		return
	}

	brothInput := orderInput.BrothId
	proteinInput := orderInput.ProteinId
	switch {
	case brothInput == "" && proteinInput == "":
		helpers.CustomError(w, "brothId and proteinId are required", http.StatusBadRequest)
		return
	case brothInput == "":
		helpers.CustomError(w, "brothId is required", http.StatusBadRequest)
		return
	case proteinInput == "":
		helpers.CustomError(w, "proteinId is required", http.StatusBadRequest)
		return
	}

	order, err := models.CreateOrder(secretKey, brothInput, proteinInput)
	if err != nil {
		errorMessage := fmt.Sprint("could not place order: ", err)
		helpers.CustomError(w, errorMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
