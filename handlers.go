package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const apiRobloxUrl string = "https://presence.roblox.com/v1/presence/users"

type RequestBody struct {
	UserIds []int64 `json:"userIds"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := RequestBody{
		UserIds: []int64{7391737385, 9158091036},
	}
	jsonBody, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.Post(apiRobloxUrl, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
