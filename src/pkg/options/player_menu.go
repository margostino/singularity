package options

import (
	"fmt"
	"org.gene/singularity/pkg/db"
)

type PlayerMenu struct {
	Id   int
	Prev int
}

func (menu PlayerMenu) show() {
	showPlayerMenu()
}

func (menu PlayerMenu) prev() int {
	return menu.Prev
}

func (menu PlayerMenu) id() int {
	return menu.Id
}

func NewPlayerMenu() PlayerMenu {
	return PlayerMenu{Id: PLAYER, Prev: MAIN}
}

func showPlayerMenu() {
	fmt.Println("---------------------")
	fmt.Println("Player Menu")
	fmt.Println("[1] - Planning")
	fmt.Println("[2] - Show stats")
	fmt.Println("[3] - Run")
	fmt.Println("[x] - Back to Main Menu")
	fmt.Println("---------------------")
}

func processPlayerOption(option int, input Input) Input {
	switch option {
	case 1:
		planning()
	case 2:
		showStats()
	case 3:
		db.CreatePlayer()
		PrintMainMenu()
	case 4:
		showPlayers()
	case 5:
		run()
	default:
		PrintMainMenu()
		input.Menu = NewMainMenu()
	}

	return input
}

func planning() {
	fmt.Println("planning")
}

func showStats() {
	fmt.Println("show stats")
}

func run() {
	fmt.Println("run")
}
