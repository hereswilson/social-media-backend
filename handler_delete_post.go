package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no id provided to handlerDeletePost"))
		return
	}
	err := apiCfg.dbClient.DeletePost(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
