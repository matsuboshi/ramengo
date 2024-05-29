package handlers

import (
	"fmt"
	"net/http"
)

func PostOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ORDERSSS!")
}
