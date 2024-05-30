package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/helpers"
	"github.com/matsuboshi/ramengo/internal/models"
)

func GetProteins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.CustomError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	proteins, err := models.GetProteins()
	if err != nil {
		helpers.CustomError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proteins)
}
