package commands

import (
	"fmt"
	"os"
)

type Exit struct {
	Action
}

func NewExitAction() *Exit {
	action := Action{
		name:  "exit",
		apply: ExecuteExit,
	}
	return &Exit{action}
}

func ExecuteExit() {
	fmt.Println("bye!")
	os.Exit(0)
}
