package option

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"org.gene/singularity/pkg/context"
)

func Welcome() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Welcome to Singularity! Please select a command.")
	fmt.Println("------------------------------------------------")
}

func Prompt() string {
	username := context.GetUsername()
	prefix := fmt.Sprintf("@%s> ", username)
	return prompt.Input(prefix, Completer)
}
