package handlers

import (
	"fmt"
	"net/http"
)

func GetBroths(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Broths!")
}
