package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handleRetrievePosts(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/posts/")
	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided to handlerRetrievePosts"))
		return
	}
	posts, err := apiCfg.dbClient.GetPosts(userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
