package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middleWareCors(mux)
	server := http.Server{
		Handler: corsMux,
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func middleWareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
