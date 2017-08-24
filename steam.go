package main

import (
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Steam struct {
	url   string
}

var spaceClient = http.Client{
	Timeout: time.Second * 2, // Maximum of 2 secs
}

func (steam *Steam) getPlayerResponse() PlayerResponse {

	req, err := http.NewRequest(http.MethodGet, steam.url, nil)
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

	response := PlayerResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("could not unmarshal json: ", err)
	}

	return response
}

func (steam *Steam) getRecentlyPlayedGames() RecentlyPlayed {
	req, err := http.NewRequest(http.MethodGet, steam.url, nil)
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

	response := RecentlyPlayed{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("could not unmarshal json: ", err)
	}

	return response
}
