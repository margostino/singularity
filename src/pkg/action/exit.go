package action

import (
	"fmt"
	"org.gene/singularity/pkg/context"
	"os"
)

type Exit struct {
	Action
}

func ExecuteExit() {
	if context.IsPlayerSelected() {
		context.Exit()
	} else {
		fmt.Println("bye!")
		os.Exit(0)
	}
}
