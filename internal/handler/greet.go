package handler

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
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
