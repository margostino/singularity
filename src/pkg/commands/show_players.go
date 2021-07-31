package commands

import "fmt"

type ShowPlayers struct {
	Action
}

func NewShowPlayersAction() *ShowPlayers {
	action := Action{
		name:  "show players",
		apply: ExecuteShowPlayer,
	}
	return &ShowPlayers{action}
}

func ExecuteShowPlayer() {
	fmt.Println("show players")
}
