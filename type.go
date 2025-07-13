package main

const (
	ERR_PLAYER_NOT_FOUND      string = "ERR_PLAYER_NOT_FOUND"
	ERR_INTERNAL_SERVER_ERROR string = "ERR_INTERNAL_SERVER_ERROR"
)

type Config struct {
	Port    int      `json:"port"`
	Bind    string   `json:"bind,omitempty"`
	Players []string `json:"players"`
}

type Player struct {
	Name  string `json:"name"`
	Score int32  `json:"score"`
}

type Players struct {
	Players []Player `json:"players"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
