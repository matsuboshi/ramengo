package main

import (
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handlers"
	"github.com/matsuboshi/ramengo/internal/middleware"
)

func greet(w http.ResponseWriter, r *http.Request) {
	defaultMessage := fmt.Sprint(
		"HELLO!!!\n",
		"YOU SHOULD TRY THESE ROUTES INSTEAD:\n",
		"- GET /broths\n",
		"- GET /proteins\n",
		"- POST /orders\n",
		"- POST /order\n",
	)

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
