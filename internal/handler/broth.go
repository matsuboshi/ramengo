package handler

import (
	"encoding/json"
	"net/http"

	customError "github.com/matsuboshi/ramengo/internal/error"
	"github.com/matsuboshi/ramengo/internal/model"
)

func ListBroths(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		customError.WriteMessage(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	broths, err := model.AllBroths()
	if err != nil {
		customError.WriteMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broths)
}
