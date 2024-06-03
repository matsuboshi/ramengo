package middleware

import (
	"fmt"
	"net/http"
	"os"

	customError "github.com/matsuboshi/ramengo/internal/error"
	"golang.org/x/crypto/bcrypt"
)

const (
	allowedOrigin  = "https://tech.redventures.com.br"
	allowedMethods = "GET, POST, OPTIONS"
	allowedHeaders = "Content-Type, x-api-key"
	credentials    = "true"
)

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Allow-Credentials", credentials)
}

func authorizeRequest(r *http.Request) error {
	storedHashedKey := os.Getenv("STORED_HASHED_KEY")
	if storedHashedKey == "" {
		return fmt.Errorf("STORED_HASHED_KEY env var not set")
	}

	incomingKey := r.Header.Get("x-api-key")
	if incomingKey == "" {
		return fmt.Errorf("x-api-key header missing")
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedHashedKey), []byte(incomingKey))
	if err != nil {
		return fmt.Errorf("api key does not match")
	}

	return nil
}

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setCorsHeaders(w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if err := authorizeRequest(r); err != nil {
			customError.WriteMessage(w, err.Error(), http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
