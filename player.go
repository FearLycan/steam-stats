package main

import "time"

//0 - Offline, 1 - Online, 2 - Busy, 3 - , 4 - Snooze, 5 - looking to trade, 6 - looking to play.

const STATUS_OFFLINE = 0
const STATUS_ONLINE = 1
const STATUS_BUSY = 2
const STATUS_AWAY = 3
const STATUS_SNOOZE = 4
const STATUS_READY_TO_TRADE = 5
const STATUS_READY_TO_PLAY = 6

type Response struct {
	Response Players
}

type Players struct {
	Players []Person
}

type Person struct {
	Name        string `json:"personaname"`
	Lastlogoff  int32 `json:"lastlogoff"`
	PersonState byte `json:"personastate"`
}

func (p *Person) lastSeen() string {
	tm := time.Unix(int64(p.Lastlogoff), 0)
	return tm.Format(time.UnixDate)

}

func (p *Person) dayAgo() int32 {
	tm := int32(time.Now().Unix()) - p.Lastlogoff
	return (tm / 60) / 24
}

func getStatusNames() map[byte]string {
	return map[byte]string{
		STATUS_OFFLINE:        "Offline",
		STATUS_ONLINE:         "Online",
		STATUS_BUSY:           "Busy",
		STATUS_AWAY:           "Away",
		STATUS_SNOOZE:         "Snooze",
		STATUS_READY_TO_TRADE: "Looking to trade",
		STATUS_READY_TO_PLAY:  "Looking to play",
	}
}

func (p *Person) getStatusName() string {
	return getStatusNames()[p.PersonState]
}
