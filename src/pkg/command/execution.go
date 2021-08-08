package command

type ExecutionPlan struct {
	Plan    []string
	Command *Command
}

func Prepare(plan []string) *ExecutionPlan {
	return &ExecutionPlan{Plan: plan}
}

func (e *ExecutionPlan) With(command *Command) *ExecutionPlan {
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
