package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/helpers"
	"github.com/matsuboshi/ramengo/internal/models"
)

type OrderInput struct {
	ProteinId string `json:"proteinId"`
	BrothId   string `json:"brothId"`
}

func (o *OrderInput) ValidateData() error {
	switch {
	case o.BrothId == "" && o.ProteinId == "":
		return errors.New("brothId and proteinId are required")
	case o.BrothId == "":
		return errors.New("brothId is required")
	case o.ProteinId == "":
		return errors.New("proteinId is required")
	default:
		return nil
	}
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

	if err := orderInput.ValidateData(); err != nil {
		helpers.CustomError(w, err.Error(), http.StatusBadRequest)
		return
	}

	secretKey := r.Header.Get("x-api-key")
	if secretKey == "" {
		helpers.CustomError(w, "x-api-key header missing", http.StatusForbidden)
		return
	}

	order, err := models.CreateOrder(secretKey, orderInput.BrothId, orderInput.ProteinId)
	if err != nil {
		errorMessage := fmt.Sprint("could not place order: ", err)
		helpers.CustomError(w, errorMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
