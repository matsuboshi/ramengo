package main

import (
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handler"
	"github.com/matsuboshi/ramengo/internal/middleware"
)

func main() {
	http.HandleFunc("/", handler.Greet)
	http.HandleFunc("/broths", middleware.Cors(handler.GetBroths))
	http.HandleFunc("/proteins", middleware.Cors(handler.GetProteins))
	http.HandleFunc("/order", middleware.Cors(handler.PostOrder))
	http.HandleFunc("/orders", middleware.Cors(handler.PostOrder))
	http.ListenAndServe(":8080", nil)
}
