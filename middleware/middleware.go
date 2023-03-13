package middleware

import (
	"net/http"

	"github.com/ShintaNakama/my-graphql-example/ctxutil"
)

func GetRoleKeyFromHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Role")
		ctx := ctxutil.SetRoleKey(r.Context(), key)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
