package commands

import (
	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
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
