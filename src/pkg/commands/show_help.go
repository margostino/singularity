package commands

import "fmt"

type ShowHelp struct {
	Action
}

func NewShowHelpAction() *ShowHelp {
	action := Action{
		name:  "show rules",
		apply: ExecuteShowHelp,
	}
	return &ShowHelp{action}
}

func ExecuteShowHelp() {
	fmt.Println("show help")
}
