package auth

import (
	"anviz-mssql-api/api/httpx"
	"net/http"
	"strings"
)

type Scope string

const (
	ReadScope Scope = "read"
	CrudScope Scope = "crud"
)

func Authorize(apiKeys *APIKeyStore, required Scope, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			httpx.WriteError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		scope, exists := apiKeys.ScopeFor(parts[1])
		if !exists {
			httpx.WriteError(w, http.StatusUnauthorized, "Invalid API key")
			return
		}

		if scope != required && scope != CrudScope {
			httpx.WriteError(w, http.StatusForbidden, "Insufficient permissions")
			return
		}

		next.ServeHTTP(w, r)
	})
}
