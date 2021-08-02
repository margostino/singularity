package main

import (
	"fmt"
	"org.gene/singularity/pkg/commands"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/preload"
	"os"
	"strings"
)

var commandTree = commands.Load()

func main() {
	commands.Welcome()
	config.LoadConfiguration()
	preload.Preload()
	plan := Input()
	Validate(plan)
	Loop(plan)
}

func Loop(plan []string) {
	for {
		action := commandTree.LookupAction(plan)
		action.Execute()
		plan = Input()
	}
}

func Validate(plan []string) {
	if !commandTree.IsValidPlan(plan) {
		fmt.Printf("command plan %q is not valid\n", plan)
		os.Exit(1)
	}
}

func Input() []string {
	commandLine := commands.Prompt()
	return strings.Fields(commandLine)
}
