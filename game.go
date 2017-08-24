package main

type RecentlyPlayed struct {
	Body struct {
		Games []RecentlyPlayedGames `json:"games"`
	} `json:"response"`
	Total byte `json:"total_count"`
}

type RecentlyPlayedGames struct {
	Name            string `json:"name"`
	Playtime2weeks  int32  `json:"playtime_2weeks"`
	PlaytimeForever int32 `json:"playtime_forever"`
}
