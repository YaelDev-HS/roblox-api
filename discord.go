package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const url string = "https://discord.com/api/webhooks/1421703040477696152/zOXt2mY0fXifUmUprOc1SfNsQ4WE90vPySqD_W5wlM-1HAvwTsEKkxzY5oICkHr5HGI1"

func sendMessage(msg string) {
	payload := map[string]string{"content": msg}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error enviando el mensaje = %s\n", err)
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("Error enviando el mensaje = %s\n", err)
		return
	}

	defer res.Body.Close()
}
