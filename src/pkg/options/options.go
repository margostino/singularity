package options

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EXIT   int = -1
	MAIN   int = 0
	PLAYER int = 1
	RUNNER int = 2
)

type Menu interface {
	id() int
	prev() int
	show()
}

type Input struct {
	Value *string
	Menu  Menu
}

func ProcessInput(reader bufio.Reader) string {
	optionText, _ := reader.ReadString('\n')
	return strings.Replace(optionText, "\n", "", -1)
}

func validateInput(err error) {
	if err != nil {
		fmt.Printf("Invalid option: %v\n", err)
		PrintMainMenu()
	}
}

func ProcessOption(input Input) Input {

	if *input.Value == "x" {
		processExitOption(input)
	}

	option, error := strconv.Atoi(*input.Value)
	validateInput(error)

	switch input.Menu.id() {
	case MAIN:
		input = processMainOption(option, input)
	case PLAYER:
		input = processPlayerOption(option, input)
	case RUNNER:
		input = processRunnerOption(option, input)
	default:
		input.Menu = nil
		exit()
	}

	return input
}

func processExitOption(input Input) Input {
	switch input.Menu.prev() {
	case MAIN:
		PrintMainMenu()
		input.Menu = NewMainMenu()
	case PLAYER:
		showPlayerMenu()
		input.Menu = NewPlayerMenu()
	default:
		input.Menu = nil
		exit()
	}
	return input
}

func exit() {
	fmt.Println("Bye!")
	os.Exit(0)
}

//func clear() {
//	cmd := exec.Command("clear")
//	cmd.Stdout = os.Stdout
//	cmd.Run()
//}
