package action

import (
	"fmt"
	"org.gene/singularity/pkg/shell"
)

func ExecuteHelp() {
	options := shell.PowerShell.GetOptions()
	for _, option := range options {
		fmt.Printf("[ %s ] - %s\n", option.Text, option.Description)
	}
}
