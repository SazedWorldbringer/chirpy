package main

import (
	"encoding/json"
	"net/http"

	"github.com/SazedWorldbringer/chirpy/internal/utils"
)

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	type chirp struct {
		Chirp string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	body := chirp{}
	err := decoder.Decode(&body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error decoding params: "+err.Error())
		return
	}

	if len(body.Chirp) > 140 {
		type ErrorResp struct {
			Error string `json:"error"`
		}
		errResp := ErrorResp{Error: "Chirp is too long"}
		utils.RespondWithJSON(w, http.StatusBadRequest, errResp)
		return
	}

	type ValidResp struct {
		Valid bool `json:"valid"`
	}
	validResp := ValidResp{Valid: true}
	utils.RespondWithJSON(w, http.StatusOK, validResp)
}
