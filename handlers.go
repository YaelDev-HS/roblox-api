package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const apiRobloxUrl string = "https://presence.roblox.com/v1/presence/users"

type UserPresenceType int8

type UserDetails struct {
	name   string
	status UserPresenceType
}

type RequestBody struct {
	UserIds []int64 `json:"userIds"`
}

type RobloxUser struct {
	UserPresenceType UserPresenceType `json:"userPresenceType"`
	UserId           int64            `json:"userId"`
}

type ResponseBody struct {
	UserPresences []RobloxUser `json:"userPresences"`
}

var memoUsers map[int64]UserDetails = make(map[int64]UserDetails)

// []UserDetails{
// 	{name: "SlayerRed_X", userId: 9385162663},
// 	{name: "132SleepingBeauty", userId: 1021391066},
// 	{name: "Lord_Belfegor", userId: 1179074352},
// 	{name: "Warden_Void", userId: 4478471035},
// 	{name: "SlimmerMain", userId: 1735183680},
// 	{name: "GetIdFromUsername", userId: 8207307371},
// 	{name: "1xjack1x2", userId: 2531463803},
// }

func start() {
	memoUsers[9158091036] = UserDetails{
		name:   "SlayerRed_X",
		status: 0,
	}
	memoUsers[1021391066] = UserDetails{
		name:   "132SleepingBeauty",
		status: 0,
	}
	memoUsers[1179074352] = UserDetails{
		name:   "Lord_Belfegor",
		status: 0,
	}
	memoUsers[4478471035] = UserDetails{
		name:   "Warden_Void",
		status: 0,
	}
	memoUsers[1735183680] = UserDetails{
		name:   "SlimmerMain",
		status: 0,
	}
	memoUsers[8207307371] = UserDetails{
		name:   "GetIdFromUsername",
		status: 0,
	}
	memoUsers[2531463803] = UserDetails{
		name:   "1xjack1x2",
		status: 0,
	}
}

func notifyChangeUserStatus(user *UserDetails) {
	log.Println("El usuario = " + user.name + " ha cambiado de status")

	switch user.status {
	case 1:
		log.Println("Esta activo")
	case 2:
		log.Println("Esta en Roblox Studio")
	case 0:
		log.Println("Se ha desconectado del juego")
	}
}

func handleChangeUserStatus(user RobloxUser) {
	player, ok := memoUsers[user.UserId]

	if !ok {
		log.Printf("User not found = %d", user.UserId)
		return
	}

	log.Println("User exists")

	if player.status != user.UserPresenceType {
		player.status = user.UserPresenceType
		memoUsers[user.UserId] = player
		notifyChangeUserStatus(&player)
	}
}

func handleUsersStatus(body ResponseBody) {
	for _, v := range body.UserPresences {
		handleChangeUserStatus(v)
		log.Printf("Viendo a = %d", v.UserId)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := RequestBody{
		UserIds: []int64{},
	}

	for i := range memoUsers {
		body.UserIds = append(body.UserIds, i)
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

	var responseData ResponseBody

	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handleUsersStatus(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success!!!"))
}
