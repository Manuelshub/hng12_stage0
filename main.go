package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")

			nowTime := time.Now().UTC()

			message := Message{
				Email:           "kcemmy03@gmail.com",
				CurrentDateTime: nowTime.Format(time.RFC3339),
				GithubUrl:       "https://github.com/Manuelshub/hng12_stage0",
			}

			err := json.NewEncoder(w).Encode(message)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
