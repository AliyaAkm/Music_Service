package middleware

import (
	"PracticeCrud/internal/domain"
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuth(repo domain.UserRepository) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Basic ") || len(auth) <= len("Basic ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
			if err != nil {
				http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
				return
			}
			pair := strings.SplitN(string(payload), ":", 2)
			if len(pair) != 2 {
				http.Error(w, "Malformed credentials", http.StatusUnauthorized)
				return
			}
			user, err := repo.FindByUsername(pair[0])
			if err != nil || user.Password != pair[1] {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
