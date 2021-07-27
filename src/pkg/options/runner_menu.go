package options

import (
	"fmt"
	"org.gene/singularity/pkg/db"
)

type RunnerMenu struct {
	Id   int
	Prev int
}

func (menu RunnerMenu) show() {
	showRunnerMenu()
}

func (menu RunnerMenu) prev() int {
	return menu.Prev
}

func (menu RunnerMenu) id() int {
	return menu.Id
}

func NewRunnerMenu() RunnerMenu {
	return RunnerMenu{Id: RUNNER, Prev: PLAYER}
}

func showRunnerMenu() {
	fmt.Println("---------------------")
	fmt.Println("Runner Menu")
	fmt.Println("[1] - Run a day")
	fmt.Println("[2] - Run a week")
	fmt.Println("[3] - Run a month")
	fmt.Println("[4] - Run a year")
	fmt.Println("[x] - Back to Player Menu")
	fmt.Println("---------------------")
}

func processRunnerOption(option int, input Input) Input {
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
