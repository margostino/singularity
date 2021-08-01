package options

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"org.gene/singularity/pkg/db"
)

type MainMenu struct {
	Id   int
	Prev int
}

func (menu MainMenu) show() {
	//PrintMainMenu()
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

}
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show help", Description: "Show Game Rules and Commands"},
		{Text: "show players", Description: "Show current players"},
		{Text: "show stats", Description: "Show game stats"},
		{Text: "create player", Description: "Create new player"},
		{Text: "start", Description: "Start the game"},
		{Text: "exit", Description: "Exit the game"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
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
