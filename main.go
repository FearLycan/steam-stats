package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

func main() {

	STEAM_ID := "76561198087405678"

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

	response := Response{}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}



	fmt.Printf("Hello %s!\n", response)
	//fmt.Printf("Last seen on Steam: %s (%d days ago) \n", response.Response.Players[0].lastSeen(), response.Response.Players[0].dayAgo())
	//fmt.Printf("Status: %s \n", response.Response.Players[0].getStatusName())
}
