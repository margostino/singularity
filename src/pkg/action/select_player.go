package action

import (
	"fmt"
)

func ExecuteSelectPlayer(args []string) {
	player := args[0]
	fmt.Printf("player %s selected\n", player)
}
