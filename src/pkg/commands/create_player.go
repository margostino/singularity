package commands

import "fmt"

type CreatePlayer struct {
	Action
}

func NewCreatePlayerAction() *CreatePlayer {
	createPlayerAction := CreatePlayer{}
	action := Action{
		name: "create player",
	}
	action.apply = (createPlayerAction).Execute
	return &CreatePlayer{action}
}

func (a CreatePlayer) Execute() {
	fmt.Println("create player")
}
