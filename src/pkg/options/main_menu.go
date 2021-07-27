package options

import (
	"fmt"
	"org.gene/singularity/pkg/db"
)

type MainMenu struct {
	Id   int
	Prev int
}

func (menu MainMenu) show() {
	PrintMainMenu()
}

func (menu MainMenu) prev() int {
	return menu.Prev
}

func (menu MainMenu) id() int {
	return menu.Id
}

func NewMainMenu() MainMenu {
	return MainMenu{Id: MAIN, Prev: EXIT}
}

func PrintMainMenu() {
	//clear()
	fmt.Println("---------------------")
	fmt.Println("Main Menu")
	fmt.Println("[1] - Show Rules")
	fmt.Println("[2] - Create Player")
	fmt.Println("[3] - Show Players")
	fmt.Println("[4] - Start")
	fmt.Println("[x] - Exit")
	fmt.Println("---------------------")
	printPrompt()
}

func printPrompt()  {
	fmt.Print("@>> ")
}
func processMainOption(option int, input Input) Input {
	switch option {
	case 1:
		showRules()
	case 2:
		db.CreatePlayer()
		PrintMainMenu()
	case 3:
		showPlayers()
	case 4:
		showPlayerMenu()
		input.Menu = NewPlayerMenu()
	default:
		PrintMainMenu()
		input.Menu = NewMainMenu()
	}

	return input
}

func showRules() {
	fmt.Println("show rules")
	PrintMainMenu()
}

func showPlayers() {
	fmt.Printf("%v", db.Players)
	PrintMainMenu()
}
