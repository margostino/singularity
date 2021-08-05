package main

import (
	"fmt"
	"org.gene/singularity/pkg/command"
	"org.gene/singularity/pkg/config"
	"org.gene/singularity/pkg/option"
	"org.gene/singularity/pkg/preload"
	"strings"
)

var commandTree = command.Load()

func main() {
	option.Welcome()
	config.LoadConfiguration()
	preload.Preload()
	plan := Input()
	Loop(plan)
}

func Loop(plan []string) {
	for {
		if Validate(plan) {
			command := commandTree.LookupCommand(plan)
			if command.Args > 0 {
				args := plan[len(plan)-command.Args:]
				command.ExecuteWith(args)
			} else {
				command.Execute()
			}
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
	commandLine := option.Prompt()
	return strings.Fields(commandLine)
}
