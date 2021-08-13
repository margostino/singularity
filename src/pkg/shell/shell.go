package shell

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/context"
	"strconv"
	"strings"
)

type Shell struct {
	Suggestions []prompt.Suggest
}

var PowerShell *Shell

func Welcome() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Welcome to Singularity! Please select a command.")
	fmt.Println("------------------------------------------------")
}

func (s *Shell) Prompt() string {
	username := context.GetUsername()
	prefix := fmt.Sprintf("@%s> ", username)
	return prompt.Input(prefix, Completer(s.Suggestions))
}

func (s *Shell) Input() []string {
	commandLine := s.Prompt()
	return strings.Fields(commandLine)
}

func NewShell() *Shell {
	var suggestions = make([]prompt.Suggest, 0)
	commands := config.GetCommandsConfiguration().CommandList

	for _, command := range commands {
		var commandText string
		if command.Args > 0 {
			commandText = command.Id + " x" + strconv.Itoa(command.Args)
		} else {
			commandText = command.Id
		}
		suggestion := prompt.Suggest{
			Text:        commandText,
			Description: command.Description,
		}
		suggestions = append(suggestions, suggestion)
	}
	PowerShell = &Shell{Suggestions: suggestions}
	return PowerShell
}

func Completer(suggestions []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func (s *Shell) GetOptions() []prompt.Suggest {
	return s.Suggestions
}
