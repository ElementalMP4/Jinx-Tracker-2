package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

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

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerName := vars["player"]

	if !contains(config.Players, playerName) {
		w.WriteHeader(http.StatusNotFound)
		errorContainer := Error{
			Message: fmt.Sprintf("Player %s not found", playerName),
			Code:    ERR_PLAYER_NOT_FOUND,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	player, err := Get(playerName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer := Error{
			Message: err.Error(),
			Code:    ERR_INTERNAL_SERVER_ERROR,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	json.NewEncoder(w).Encode(player)
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

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
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

	sort.Slice(players, func(i, j int) bool {
		return players[i].Score > players[j].Score
	})

	renderTemplate(w, "game", players)
}

func incrementPlayerRoute(w http.ResponseWriter, r *http.Request) {
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

	_, err := Increment(player)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer := Error{
			Message: err.Error(),
			Code:    ERR_INTERNAL_SERVER_ERROR,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func decrementPlayerRoute(w http.ResponseWriter, r *http.Request) {
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

	_, err := Decrement(player)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContainer := Error{
			Message: err.Error(),
			Code:    ERR_INTERNAL_SERVER_ERROR,
		}
		json.NewEncoder(w).Encode(errorContainer)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
