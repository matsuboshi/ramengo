package middleware

import (
	"net/http"

	"github.com/matsuboshi/ramengo/internal/helpers"
	"golang.org/x/crypto/bcrypt"
)

var storedHashedKey = "$2a$10$oX/kWTonHw6YZm6apfplO.XX0ZxRAewn845gBWGMGVp/wMrgDQJwK"

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://tech.redventures.com.br")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		incomingKey := r.Header.Get("x-api-key")

		if incomingKey == "" {
			helpers.CustomError(w, "x-api-key header missing", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(storedHashedKey), []byte(incomingKey))
		if err != nil {
			helpers.CustomError(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
