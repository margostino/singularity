package main

import (
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jasonlvhit/gocron"
	"org.gene/singularity/pkg/commands"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/options"
	"org.gene/singularity/pkg/preload"
	"os"
	"strings"
	"time"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show help", Description: "Show Game Rules and Commands"},
		{Text: "show players", Description: "Show current players"},
		{Text: "show stats", Description: "Show game stats"},
		{Text: "create player", Description: "Create new player"},
		{Text: "start", Description: "Start the game"},
		{Text: "exit", Description: "Exit the game"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	//go executeCronJob()
	options.Welcome()
	commandTree := commands.Load()
	config.LoadConfiguration()
	preload.Preload()
	plan := Input()
	Validate(plan, commandTree)
	Loop(plan, commandTree)
}

func Loop(plan []string, commandTree *commands.CommandTree) {
	for {
		action := commandTree.Process(plan)
		action.Execute()
		plan = Input()
	}
}

func Validate(plan []string, commandTree *commands.CommandTree) {
	if !commandTree.IsValidPlan(plan) {
		fmt.Printf("command plan %q is not valid\n", plan)
		os.Exit(1)
	}
}

func Input() []string {
	commandLine := options.Prompt()
	return strings.Fields(commandLine)
}

func process() {
	var input = options.Input{Value: nil, Menu: options.NewMainMenu()}
	reader := bufio.NewReader(os.Stdin)
	for {
		option := options.ProcessInput(*reader)
		input.Value = &option
		input = options.ProcessOption(input)
	}
}

func myTask() {
	fmt.Println("This task will run periodically")
}
func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}
