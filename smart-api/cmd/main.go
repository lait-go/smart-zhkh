package main

import (
	"net/http"
	"smart-api/api"
	"smart-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", api.RegisterHandler)
	mux.HandleFunc("/api/login", api.LoginHandler)
	mux.HandleFunc("/api/charges", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargesHandler(w, r)
		} else if r.Method == http.MethodPost {
			api.CreateChargeHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}) 

	handler := middleware.CorsMiddleware(mux)
	http.ListenAndServe(":8080", handler)
}