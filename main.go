package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/hereswilson/social-media-backend/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func main() {
	mux := http.NewServeMux()

	client := database.NewClient("db.json")

	err := client.EnsureDatabase()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		dbClient: client,
	}

	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/err", testErrHandler)
	mux.HandleFunc("/users", apiCfg.endpointUsersHandler)
	mux.HandleFunc("/users/", apiCfg.endpointUsersHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Println("Starting server at", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

type errorBody struct {
	Error string `json:"error"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// you can use any compatible type, but let's use our database package's User type for practice
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("test error"))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
