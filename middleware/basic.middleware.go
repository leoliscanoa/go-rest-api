package middleware

import (
	"net/http"
)

func BasicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user, pwd, ok := request.BasicAuth()
		if !ok || user != "admin" || pwd != "admin" {
			writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
