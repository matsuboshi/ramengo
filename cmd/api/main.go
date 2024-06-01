package main

import (
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handler"
	"github.com/matsuboshi/ramengo/internal/middleware"
)

func main() {
	http.HandleFunc("/", handler.Greet)
	http.HandleFunc("/broths", middleware.Cors(handler.ListBroths))
	http.HandleFunc("/proteins", middleware.Cors(handler.ListProteins))
	http.HandleFunc("/order", middleware.Cors(handler.PostOrder))
	http.HandleFunc("/orders", middleware.Cors(handler.PostOrder))
	http.ListenAndServe(":8080", nil)
}
