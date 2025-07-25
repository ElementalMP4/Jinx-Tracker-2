package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	config *Config
	err    error
)

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decoding config: %w", err)
	}

	return &cfg, nil
}

func main() {
	config, err = LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	addr := fmt.Sprintf("%s:%d", config.Bind, config.Port)

	fmt.Printf("Players: %s\n", config.Players)

	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/api/players", getPlayersHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/increment/{player}", incrementPlayerHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/decrement/{player}", decrementPlayerHandler).Methods(http.MethodPost)

	// UI Routes
	r.HandleFunc("/", indexPageHandler)
	r.HandleFunc("/increment/{player}", incrementPlayerRoute).Methods(http.MethodPost)
	r.HandleFunc("/decrement/{player}", decrementPlayerRoute).Methods(http.MethodPost)

	// Static server
	r.PathPrefix("/").Handler(http.FileServer(http.FS(staticFS)))

	fmt.Println("Starting server on", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
