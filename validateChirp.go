package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SazedWorldbringer/chirpy/internal/utils"
)

func rmProfane(chirp string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Fields(chirp)

	wordsToReplace := make(map[string]string)
	for _, word := range profaneWords {
		wordsToReplace[word] = "****"
	}

	modifiedWords := []string{}
	for _, word := range words {
		lowercaseWord := strings.ToLower(word)
		if replacement, ok := wordsToReplace[lowercaseWord]; ok {
			modifiedWords = append(modifiedWords, replacement)
		} else {
			modifiedWords = append(modifiedWords, word)
		}
	}

	return strings.Join(modifiedWords, " ")
}

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

	cleanedChirp := rmProfane(body.Chirp)

	type CleanedResp struct {
		CleanedBody string `json:"cleaned_body"`
	}
	validResp := CleanedResp{CleanedBody: cleanedChirp}
	utils.RespondWithJSON(w, http.StatusOK, validResp)
}
