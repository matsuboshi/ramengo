package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	customError "github.com/matsuboshi/ramengo/internal/error"
	"github.com/matsuboshi/ramengo/internal/repository"
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
		customError.WriteMessage(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var orderInput OrderInput

	if err := json.NewDecoder(r.Body).Decode(&orderInput); err != nil {
		errorMessage := fmt.Sprintf("error decoding order input: %v", err)
		customError.WriteMessage(w, errorMessage, http.StatusBadRequest)
		return
	}

	if err := orderInput.ValidateData(); err != nil {
		customError.WriteMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	secretKey := r.Header.Get("x-api-key")
	if secretKey == "" {
		customError.WriteMessage(w, "x-api-key header missing", http.StatusForbidden)
		return
	}

	order, err := repository.CreateOrder(secretKey, orderInput.BrothId, orderInput.ProteinId)
	if err != nil {
		errorMessage := fmt.Sprint("could not place order: ", err)
		customError.WriteMessage(w, errorMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
