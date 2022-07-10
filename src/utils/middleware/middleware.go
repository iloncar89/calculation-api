package middleware

import "net/http"

//Middleware function is used to add Access-Control-Allow-Origin, Content-Type, and Access-Control-Allow-Methods in response header.
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		handler.ServeHTTP(w, r)
	})
}
