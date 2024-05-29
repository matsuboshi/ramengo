package main

import (
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handlers"
)

const defaultMessage = `YOU SHOULD TRY THESE ROUTES INSTEAD:
- GET /broths
- GET /proteins
- POST /orders`

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, defaultMessage)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/broths", handlers.GetBroths)
	http.HandleFunc("/proteins", handlers.GetProteins)
	http.HandleFunc("POST /orders", handlers.PostOrder)
	http.ListenAndServe(":8080", nil)
}
