package main

import (
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handlers"
	"github.com/matsuboshi/ramengo/internal/middleware"
)

const defaultMessage = `YOU SHOULD TRY THESE ROUTES INSTEAD:
- GET /broths
- GET /proteins
- POST /orders
- POST /order`

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, defaultMessage)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/broths", middleware.CorsMiddleware(handlers.GetBroths))
	http.HandleFunc("/proteins", middleware.CorsMiddleware(handlers.GetProteins))
	http.HandleFunc("/order", middleware.CorsMiddleware(handlers.PostOrder))
	http.HandleFunc("/orders", middleware.CorsMiddleware(handlers.PostOrder))
	http.ListenAndServe(":8080", nil)
}
