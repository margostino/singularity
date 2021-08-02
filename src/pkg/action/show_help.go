package action

import (
	"fmt"
	"org.gene/singularity/pkg/option"
)

func ExecuteShowHelp() {
	options := option.GetOptions()
	for _, option := range options {
		fmt.Printf("[ %s ] - %s\n", option.Text, option.Description)
	}
}
