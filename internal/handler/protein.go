package handler

import (
	"encoding/json"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/errormsg"
	"github.com/matsuboshi/ramengo/internal/model"
)

func GetProteins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errormsg.CustomError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	proteins, err := model.GetProteins()
	if err != nil {
		errormsg.CustomError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proteins)
}