package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const url string = "https://discord.com/api/webhooks/1422026393910251571/1Ute0tXo-yuS7zEH4_8uv7ZpO4BShhrIzuDue9dcgz7PXCZR3Qyeg_1MIddoU_F9vPq9"

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
