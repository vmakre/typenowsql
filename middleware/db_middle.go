package middleware

import (
	"app_chi_templ/dbmanager"
	"context"
	"net/http"
)

type contextKey string

const dbManagerKey contextKey = "dbManager"

// DBMiddleware injects the database manager into the request context
func DBMiddleware(dbManager *dbmanager.DatabaseManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dbManagerKey, dbManager)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetDBManagerFromContext retrieves the database manager from context
func GetDBManagerFromContext(ctx context.Context) *dbmanager.DatabaseManager {
	if dbm, ok := ctx.Value(dbManagerKey).(*dbmanager.DatabaseManager); ok {
		return dbm
	}
	return nil
}
