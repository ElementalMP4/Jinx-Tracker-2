package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func createPlayerComponent(idx int, player Player) ReceiptComponent {
	return ReceiptComponent{
		Type:     "text",
		Content:  fmt.Sprintf("%d. %s: %d", (idx + 1), player.Name, player.Score),
		FontSize: "20",
		Align:    "left",
	}
}

func createTitleComponent() ReceiptComponent {
	return ReceiptComponent{
		Type:     "header",
		Content:  "Leaderboard",
		Align:    "center",
		FontSize: "fit",
		Bold:     true,
	}
}

func sendToPrinter(receipt Receipt) error {
	j, err := json.Marshal(receipt.Components)
	if err != nil {
		return err
	}

	var url = config.Printer
	if url == "" {
		return fmt.Errorf("print server URL not set")
	}

	resp, err := http.Post(url+"/print-receipt", "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("print failed: %s", body)
	}
	return nil
}
