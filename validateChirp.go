package main

import (
	"encoding/json"
	"net/http"
)

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	type chirp struct {
		Chirp string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	body := chirp{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Error decoding params: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(body.Chirp) > 140 {
		type ErrorResp struct {
			Error string `json:"error"`
		}
		errResp := ErrorResp{Error: "Chirp is too long"}
		jsonResp, err := json.Marshal(errResp)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	type ValidResp struct {
		Valid bool `json:"valid"`
	}
	validResp := ValidResp{Valid: true}
	jsonResp, err := json.Marshal(validResp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
