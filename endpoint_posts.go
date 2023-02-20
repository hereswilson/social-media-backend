package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handleRetrievePosts(w, r)
	case http.MethodPost:
		apiCfg.handleCreatePost(w, r)
	case http.MethodDelete:
		apiCfg.handleDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
