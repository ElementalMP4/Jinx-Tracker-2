package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func getPlayersHandler(w http.ResponseWriter, r *http.Request) {
	players := []Player{}
	for _, name := range config.Players {
		player, err := Get(name)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorContainer := Error{
				Message: err.Error(),
				Code:    ERR_INTERNAL_SERVER_ERROR,
			}
			json.NewEncoder(w).Encode(errorContainer)
			return
		}

		players = append(players, *player)
	}
	result := Players{
		Players: players,
	}
	json.NewEncoder(w).Encode(result)
}

func incrementPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]

	if !contains(config.Players, player) {
		w.WriteHeader(http.StatusNotFound)
		errorContainer := Error{
			Message: fmt.Sprintf("Player %s not found", player),
			Code:    ERR_PLAYER_NOT_FOUND,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	updatedPlayer, err := Increment(player)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer := Error{
			Message: err.Error(),
			Code:    ERR_INTERNAL_SERVER_ERROR,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	json.NewEncoder(w).Encode(updatedPlayer)
}

func decrementPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]

	if !contains(config.Players, player) {
		w.WriteHeader(http.StatusNotFound)
		errorContainer := Error{
			Message: fmt.Sprintf("Player %s not found", player),
			Code:    ERR_PLAYER_NOT_FOUND,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	updatedPlayer, err := Decrement(player)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer := Error{
			Message: err.Error(),
			Code:    ERR_INTERNAL_SERVER_ERROR,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	json.NewEncoder(w).Encode(updatedPlayer)
}
