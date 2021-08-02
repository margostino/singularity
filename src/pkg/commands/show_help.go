package commands

import "fmt"

func ExecuteShowHelp() {
	options := GetOptions()
	for _, option := range options {
		fmt.Printf("[ %s ] - %s\n", option.Text, option.Description)
	}
}
