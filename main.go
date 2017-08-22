package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	var (
		STEAM_ID  string = os.Getenv("STEAM_ID")
		STEAM_KEY string = os.Getenv("STEAM_KEY")
	)

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", STEAM_KEY, STEAM_ID)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("could not establish new request: ", err)
	}

	res, err := spaceClient.Do(req)
	if err != nil {
		log.Fatal("could not execute new request: ", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("could not read the body: ", err)
	}

	response := Response{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("could not unmarshal json: ", err)
	}

	fmt.Printf("Hello %#v!\n", response)
	//fmt.Printf("Last seen on Steam: %s (%d days ago) \n", response.Response.Players[0].lastSeen(), response.Response.Players[0].dayAgo())
	//fmt.Printf("Status: %s \n", response.Response.Players[0].getStatusName())
}
