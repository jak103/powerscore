package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

var currentScore = Score{
	Home: 0,
	Away: 0,
}

type Shots struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

var currentShots = Shots{
	Home: 0,
	Away: 0,
}

var scoreMutex sync.Mutex
var shotMutex sync.Mutex

const port string = "http://localhost:5174" //This line needs to be adjusted depending on what port the VUE app is running on

func getScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", port)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	scoreMutex.Lock()
	defer scoreMutex.Unlock()
	score := currentScore

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(score); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to encode score: %v", err)
		return
	}
}

func updateScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", port)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Home *int `json:"home"`
		Away *int `json:"away"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	scoreMutex.Lock()
	if requestBody.Home != nil {
		currentScore.Home = *requestBody.Home
	}
	if requestBody.Away != nil {
		currentScore.Away = *requestBody.Away
	}
	scoreMutex.Unlock()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Score updated successfully"))
}

func getShots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", port)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	shotMutex.Lock()
	defer shotMutex.Unlock()
	shots := currentShots

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(shots); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to encode score: %v", err)
		return
	}
}

func updateShots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", port)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Home *int `json:"home"`
		Away *int `json:"away"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shotMutex.Lock()
	if requestBody.Home != nil {
		currentShots.Home = *requestBody.Home
	}
	if requestBody.Away != nil {
		currentShots.Away = *requestBody.Away
	}
	shotMutex.Unlock()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shots updated successfully"))
}

func main() {
	http.HandleFunc("/api/score", getScore)
	http.HandleFunc("/api/update-score", updateScore)
	http.HandleFunc("/api/shots", getShots)
	http.HandleFunc("/api/update-shots", updateShots)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	<-stop
}
