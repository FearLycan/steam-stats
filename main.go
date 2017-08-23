package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
)

const EXIT = 10

func main() {

	var (
		STEAM_ID  string = os.Getenv("STEAM_ID")
		STEAM_KEY string = os.Getenv("STEAM_KEY")
	)

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", STEAM_KEY, STEAM_ID)

	steam := Steam{url, "player"}

	playerResponse := steam.getPlayerResponse()

	n := 0
	clearScreen();
	for n != EXIT {

		menu(playerResponse);
		fmt.Print("Enter: ")
		fmt.Scanf("%d", &n)
		clearScreen();


		switch n {
		case 1:
			fmt.Println("1!!")
		default:
			fmt.Println("What You Wanna Do?")
		}

		clearScreen();

	}

	//menu(playerResponse);

	//fmt.Printf("Hello %#v!\n", response.Body.Players[0].Name)
	//fmt.Printf("Last seen on Steam: %s (%d days ago) \n", response.Body.Players[0].lastSeen(), response.Body.Players[0].dayAgo())
	//fmt.Printf("Status: %s \n", response.Body.Players[0].getStatusName())

}

func menu(response PlayerResponse) {
	fmt.Println(response.Body.Players[0].Name, " - ", response.Body.Players[0].getStatusName())
	fmt.Println("=================================================")
	fmt.Printf("%d. \n", 1)
	fmt.Printf("%d. \n", 2)
	fmt.Printf("%d. \n", 3)
	fmt.Printf("%d. Exit \n", EXIT)
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
