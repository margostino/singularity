package runner

import (
	"fmt"
	"org.gene/singularity/pkg/command"
	"org.gene/singularity/pkg/context"
	"org.gene/singularity/pkg/shell"
)

var commandMap *command.CommandMap

type ExecutionPlan struct {
	Plan    []string
	Command *command.Command
}

func Loop() {
	var plan []string
	powershell := shell.NewShell()
	commandMap = command.Load()
	context.NewWorldContext()
	for {
		plan = powershell.Input()
		if Validate(plan) {
			command := commandMap.LookupCommand(plan)
			Prepare(plan).With(command).Execute()
		}
	}
}

func Validate(plan []string) bool {
	if len(plan) == 0 {
		return false
	}
	if !commandMap.IsValidPlan(plan) {
		fmt.Printf("command plan %q is not valid\n", plan)
		return false
	}
	return true
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
