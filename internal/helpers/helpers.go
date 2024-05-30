package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func CustomError(w http.ResponseWriter, httpStatus int, errorMsg string) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ErrorResponse{errorMsg})
}
