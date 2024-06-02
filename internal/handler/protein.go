package handler

import (
	"encoding/json"
	"net/http"

	customError "github.com/matsuboshi/ramengo/internal/error"
	"github.com/matsuboshi/ramengo/internal/repository"
)

func ListProteins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		customError.WriteMessage(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	proteins, err := repository.AllProteins()
	if err != nil {
		customError.WriteMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proteins)
}
