package action

import (
	"fmt"
	"os"
)

type Exit struct {
	Action
}

func ExecuteExit() {
	fmt.Println("bye!")
	os.Exit(0)
}
