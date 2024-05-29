package handlers

import (
	"fmt"
	"net/http"
)

func GetProteins(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Proteins!")
}
