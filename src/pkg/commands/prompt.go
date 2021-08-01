package commands

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

func Welcome() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Welcome to Singularity! Please select a command.")
	fmt.Println("------------------------------------------------")
}

func Prompt() string {
	//clear()
	return prompt.Input("@> ", Completer)
}
