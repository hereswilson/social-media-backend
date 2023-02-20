package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handleGetUser(w, r)
	case http.MethodPost:
		apiCfg.handleCreateUser(w, r)
	case http.MethodPut:
		apiCfg.handleUpdateUser(w, r)
	case http.MethodDelete:
		apiCfg.handleDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
