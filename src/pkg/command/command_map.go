package command

type CommandMap struct {
	commands map[string]*Command
}

func NewCommandMap(commands map[string]*Command) *CommandMap {
	return &CommandMap{commands}
}

func (m *CommandMap) IsValidPlan(plan []string) bool {
	if len(plan) == 0 {
		return false
	}
	command, ok := m.commands[plan[0]]
	return ok && command.Validate(plan)
}

func (m *CommandMap) LookupCommand(plan []string) *Command {
	command, _ := m.commands[plan[0]]
	return command.GetLastCommand(plan)
}
