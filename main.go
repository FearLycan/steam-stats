package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	"bufio"
)

const MENU_EXIT = 10

func main() {

	var (
		STEAM_ID  string = os.Getenv("STEAM_ID")
		STEAM_KEY string = os.Getenv("STEAM_KEY")
	)

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", STEAM_KEY, STEAM_ID)

	steam := Steam{url}

	playerResponse := steam.getPlayerResponse()

	n := 0
	clearScreen();
	for n != MENU_EXIT {

		menu(playerResponse);
		fmt.Print("Enter: ")
		fmt.Scanf("%d", &n)

		clearScreen();
		switch n {
		case 1:
			fmt.Printf("===== Recently Played Games by %s ===== \n", playerResponse.Body.Players[0].Name)

			url = fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=%s&steamid=%s&format=json", STEAM_KEY, STEAM_ID)

			steam = Steam{url}

			recentlyPlayedGames := steam.getRecentlyPlayedGames()

			for i, game := range recentlyPlayedGames.Body.Games {
				fmt.Printf("%d. %s\t played by %d\n", i, game.Name, game.Playtime2weeks)
			}

			fmt.Print("\nPress 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

		default:
			fmt.Println("What You Wanna Do?")
		}

		clearScreen();
	}

}

func menu(response PlayerResponse) {
	fmt.Println(response.Body.Players[0].Name, " 			Status: ", response.Body.Players[0].getStatusName())
	fmt.Println("=================================================")
	fmt.Printf("%d. Recently Played Games \n", 1)
	fmt.Printf("%d. \n", 2)
	fmt.Printf("%d. \n", 3)
	fmt.Printf("%d. Exit \n", MENU_EXIT)
	fmt.Println("=================================================")
}

func clearScreen() {
	switch system := runtime.GOOS; system {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		//cmd := exec.Command("cls")
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Printf("%s.\n", system)
		fmt.Println("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
