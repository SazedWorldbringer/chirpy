package main

import (
	"net/http"

	"github.com/SazedWorldbringer/chirpy/internal/utils"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
