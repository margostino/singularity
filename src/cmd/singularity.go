package main

import (
	"fmt"
	"org.gene/singularity/pkg/commands"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/preload"
	"strings"
)

var commandTree = commands.Load()

func main() {
	commands.Welcome()
	config.LoadConfiguration()
	preload.Preload()
	plan := Input()
	Loop(plan)
}

func Loop(plan []string) {
	for {
		if Validate(plan) {
			action := commandTree.LookupAction(plan)
			action.Execute()
		}
		plan = Input()
	}
}

func Validate(plan []string) bool {
	if !commandTree.IsValidPlan(plan) {
		fmt.Printf("command plan %q is not valid\n", plan)
		return false
	}
	return true
}

func Input() []string {
	commandLine := commands.Prompt()
	return strings.Fields(commandLine)
}
