package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Response struct {
    Email           string `json:"email"`
    CurrentDateTime string `json:"current_datetime"`
    GitHubURL       string `json:"github_url"`
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handler).Methods(http.MethodGet)
    logger := handlers.CombinedLoggingHandler(os.Stdout, r)

		cors := handlers.CORS(
			handlers.AllowedHeaders([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"https://*"}),
	)


    log.Fatal(http.ListenAndServe(":8000", cors(logger)))
}

func handler(w http.ResponseWriter, r *http.Request) {
    res := Response{
        Email:           "oguejioforrichard@gmail.com",
        CurrentDateTime: time.Now().UTC().Format(time.RFC3339),
        GitHubURL:       "https://github.com/ikennarichard/hng-12/",
    }

    if err := json.NewEncoder(w).Encode(res); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}