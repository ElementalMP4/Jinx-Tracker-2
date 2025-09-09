package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	device "github.com/ElementalMP4/go-hd44780"
	"github.com/ElementalMP4/go-i2c"
)

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (p Player) toString() string {
	str := fmt.Sprintf("%s: %d", p.Name, p.Score)
	padding := (16 - len(str)) / 2
	if padding > 0 {
		str = fmt.Sprintf("%s%s", spaces(padding), str)
	}
	return str
}

func spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

type Players struct {
	Players []Player `json:"players"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPlayers() (*Players, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/players", os.Getenv("JT_URL")))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var players Players
	err = json.Unmarshal(body, &players)
	if err != nil {
		return nil, err
	}

	return &players, nil
}

func initI2c() *i2c.I2C {
	i2c, err := i2c.NewI2C(0x27, 1)
	checkError(err)
	return i2c
}

func initLcd(i2c *i2c.I2C) *device.Lcd {
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	checkError(err)
	return lcd
}

func main() {
	i2c := initI2c()
	defer i2c.Close()

	lcd := initLcd(i2c)
	lcd.BacklightOn()
	lcd.Clear()

	var playerCache *Players

	for {
		players, err := getPlayers()
		if err == nil {
			if playerCache == nil || !reflect.DeepEqual(players, playerCache) {
				lcd.SetPosition(0, 0)
				fmt.Fprint(lcd, players.Players[0].toString())

				lcd.SetPosition(1, 1)
				fmt.Fprint(lcd, players.Players[1].toString())

				playerCache = players

				for i := 0; i < 5; i++ {
					lcd.BacklightOff()
					time.Sleep(100)
					lcd.BacklightOn()
					time.Sleep(100)
				}
			}
		} else {
			fmt.Printf("Failed to get players: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}
