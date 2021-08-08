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

type ExecutionPlan struct {
	Plan    []string
	Command *command.Command
}

func main() {
	option.Welcome()
	config.LoadConfiguration()
	preload.Preload()
	Loop()
}

func Loop() {
	var plan []string
	for {
		plan = Input()
		if Validate(plan) {
			command := commandTree.LookupCommand(plan)
			Prepare(plan).With(command).Execute()
		}
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

func Prepare(plan []string) *ExecutionPlan {
	// Potentially pre-processing
	return &ExecutionPlan{Plan: plan}
}

func (e *ExecutionPlan) With(command *command.Command) *ExecutionPlan {
	e.Command = command
	return e
}

func (e *ExecutionPlan) Execute() {
	if e.Command.Args > 0 {
		args := e.Plan[len(e.Plan)-e.Command.Args:]
		e.Command.ExecuteWith(args)
	} else {
		e.Command.Execute()
	}
}