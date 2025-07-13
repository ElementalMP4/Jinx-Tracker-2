package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

var gameState map[string]int32

const savePath = "game.bin"

func init() {
	gameState = make(map[string]int32)

	if _, err := os.Stat(savePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No save file found, starting new game state.")
	} else {
		file, err := os.Open(savePath)
		if err != nil {
			fmt.Println("Error opening save file:", err)
			return
		}
		defer file.Close()

		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&gameState)
		if err != nil {
			fmt.Println("Error decoding save file:", err)
			gameState = make(map[string]int32)
		} else {
			fmt.Println("Game state loaded.")
		}
	}
}

func Increment(name string) (*Player, error) {
	value := gameState[name]
	gameState[name] = value + 1

	err := save()
	if err != nil {
		return nil, err
	}

	return Get(name)
}

func Decrement(name string) (*Player, error) {
	value := gameState[name]
	if value > 0 {
		gameState[name] = value - 1
	}

	err := save()
	if err != nil {
		return nil, err
	}

	return Get(name)
}

func Get(name string) (*Player, error) {
	score, exists := gameState[name]
	if !exists {
		score = 0
		gameState[name] = score

		err := save()
		if err != nil {
			return nil, err
		}
	}

	player := Player{
		Name:  name,
		Score: score,
	}

	return &player, nil
}

func save() error {
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(gameState)
}
