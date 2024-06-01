package main

import (
	"fmt"
	"net/http"

	"github.com/matsuboshi/ramengo/internal/handler"
	"github.com/matsuboshi/ramengo/internal/middleware"
)

func main() {
	fmt.Println("")
	fmt.Println("Test the endpoint at https://tech.redventures.com.br")
	fmt.Println("API URL: http://localhost:8080")
	fmt.Println("API KEY: ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf")

	http.HandleFunc("/", handler.Greet)
	http.HandleFunc("/broths", middleware.Cors(handler.ListBroths))
	http.HandleFunc("/proteins", middleware.Cors(handler.ListProteins))
	http.HandleFunc("/order", middleware.Cors(handler.PostOrder))
	http.HandleFunc("/orders", middleware.Cors(handler.PostOrder))
	http.ListenAndServe(":8080", nil)
}
