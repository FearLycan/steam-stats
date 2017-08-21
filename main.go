package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

type response struct {
	Response players
}

type players struct {
	Players []Person
}

type Person struct {
	Name       string `json:"personaname"`
	Lastlogoff int32 `json:"lastlogoff"`
}

//0 - Offline, 1 - Online, 2 - Busy, 3 - Away, 4 - Snooze, 5 - looking to trade, 6 - looking to play.

const STATUS_OFFLINE = 0
const STATUS_ONLINE = 1
const STATUS_BUSY = 2
const STATUS_AWAY = 3
const STATUS_SNOOZE = 4
const STATUS_READY_TO_TRADE = 5
const STATUS_READY_TO_PLAY = 6

func main() {

	STEAM_ID := "#"

	KEY := "#"

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", KEY, STEAM_ID)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := response{}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	lastLogin := response.Response.Players[0].Lastlogoff

	fmt.Printf("Hello %s!\n", response.Response.Players[0].Name)
	fmt.Printf("Last seen on Steam: %s (%d days ago)", lastSeen(lastLogin), dayAgo(lastLogin))

}

func lastSeen(seconds int32) (string) {
	tm := time.Unix(int64(seconds), 0)
	return tm.Format(time.UnixDate)
}

func dayAgo(seconds int32) (int32) {
	tm := int32(time.Now().Unix()) - seconds
	return (tm / 60) / 24
}
